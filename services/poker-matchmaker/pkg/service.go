package pkg

import (
	"fmt"
	"net/http"

	"agones.dev/agones/pkg/client/clientset/versioned"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsMW "github.com/slok/go-http-metrics/middleware"
	metricsNegroni "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
	"k8s.io/client-go/rest"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/serviceconfig"
)

//Start the poker matchmaker service
func Start(logger *zap.SugaredLogger) {

	//################## Setup ####################
	//Configure Poker Classes
	classes := []models.Class{*models.NewClass(50, 999, 20),
		*models.NewClass(1000, 9999, 20),
		*models.NewClass(10000, 49999, 20),
		*models.NewClass(50000, 1000000, 20),
	}


	logger.Debugf("Node Addresses %v", viper.GetStringSlice(serviceconfig.NodeIPAddresses))

	//Connect to Kubernetes and Agones API
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {
		logger.Fatalw("Could not get cluster configuration")
	}

	agonesClient, _ := versioned.NewForConfig(kubeConfig)
	if err != nil {
		logger.Fatalw("Could not connect to agones api")
	}

	//Create poker gameserver manager
	manager := lobby.NewManager(logger, agonesClient, classes)

	//############# HTTP ################
	//Create HTTP handler
	ticketHandler := handlers.NewTicket(logger, agonesClient, manager)

	//Setup routes
	r := mux.NewRouter()
	r.Handle("/api/poker/ticket", mw.RequireAuth(ticketHandler.GetTicketWithParams)).Methods(http.MethodPost).Queries("class", "{class:[0-9]+}", "buyIn", "{buyIn:[0-9]+}")
	r.Handle("/api/poker/ticket/{id}", mw.RequireAuth(ticketHandler.GetTicketWithID)).Methods(http.MethodPost).Queries("class", "{class:[0-9]+}", "buyIn", "{buyIn:[0-9]+}")
	r.HandleFunc("/api/poker/pokerinfo", ticketHandler.PokerInfo).Methods(http.MethodGet)
	r.HandleFunc("/api/poker/classinfo", ticketHandler.ClassInfo).Methods(http.MethodGet)
	r.Handle("/metrics", promhttp.Handler())

	//Register Middleware
	metricsMiddleware := metricsMW.New(metricsMW.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
		Service:  "PokerMatchMaker",
	})

	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery(), metricsNegroni.Handler("", metricsMiddleware))
	n.UseHandler(r)

	//Configure HTTP server
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", viper.GetInt(config.HTTPPort)),
		Handler:           n,
		WriteTimeout:      viper.GetDuration(config.WriteTimeout),
		ReadHeaderTimeout: viper.GetDuration(config.ReadTimeout),
		IdleTimeout:       viper.GetDuration(config.IdleTimeout),
	}

	//Start HTTP server
	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}
