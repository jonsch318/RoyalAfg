package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"golang.org/x/net/context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/gameServer"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
)

//main method is the entry point of the game server
func main() {
	//Register Logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Configure
	config.ReadStandardConfig("poker", logger)
	serviceconfig.SetDefaults()

	//Bind environment variables
	viper.SetEnvPrefix("poker")
	_ = viper.BindEnv(config.RabbitMQUsername)
	_ = viper.BindEnv(config.RabbitMQPassword)

	//connect to rabbitmq to send user commands
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConn, err := rabbit.NewRabbitMessageBroker(logger, rabbitURL)
	if err != nil {
		logger.Fatalw("Could not connect to service bus", "error", err)
	}
	defer rabbitConn.Close()

	//Register stop signal
	go gameServer.DoSignal()

	//Creating agones sdk instance to communicate with the game server orchestrator
	logger.Infof("Creating SDK instance")
	s, err := sdk.NewSDK()
	if err != nil {
		logger.Fatalf("Error during sdk connection, %v", err)
	}

	//StartHealth ping to agones.
	logger.Info("Health Ping to agones server management")
	stop := make(chan struct{})
	go gameServer.DoHealthPing(s, stop)


	//Configure logging
	lobbyConfigured := false
	shutDownStop := make(chan interface{})
	b := bank.NewBank(rabbitConn)
	lobbyInstance := lobby.NewLobby(b, s)

	//Watch Pod Labels for lobby information
	err = s.WatchGameServer(func(gs *coresdk.GameServer) {
		if !lobbyConfigured {
			err2 := SetLobby(b, lobbyInstance, gs)
			if err2 == nil {
				logger.Warnw("Lobby configured", "id", lobbyInstance.LobbyID)
				if lobbyInstance.Count() <= 0 {
					go StartShutdownTimer(shutDownStop, s)
				}
				//Lobby is configured through kubernetes labels and is assigned a unique id.
				lobbyConfigured = true
			} else if !strings.HasPrefix(err.Error(), "key needed") {
				logger.Errorw("Error during configuration", "error", err)
			}
		}
	})
	if err != nil {
		logger.Fatalf("Error during sdk annotation subscription: %s", err)
	}

	//Register HTTP handlers
	gameHandler := handlers.NewGame(lobbyInstance, s, shutDownStop)

	//Setup Routes
	r := mux.NewRouter()
	r.HandleFunc("/api/poker/health", gameHandler.Health).Methods(http.MethodGet)
	r.HandleFunc("/api/poker/join", gameHandler.Join) //Websocket join

	//RegisterMiddleware
	n := negroni.New(negroni.NewRecovery(), mw.NewLogger(logger.Desugar()))
	n.UseHandler(r)

	//Call Ready
	err = s.Ready()
	if err != nil {
		logger.Errorw("Error during sdk ready call", "error", err)
	}
	logger.Info("SDK Ready called")


	//Configure HTTP server
	port := viper.GetString(config.HTTPPort)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
		Handler: n,
	}

	//Start HTTP server
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))


}

//StartShutdownTimer sets the status shutdown after 10 minutes.
func StartShutdownTimer(stop chan interface{}, s *sdk.SDK) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Minute))
	select {
	case _ = <-stop:
		cancel()
		//Cancel shutdown
	case <-ctx.Done():
		_ = s.Shutdown()
	}
}
