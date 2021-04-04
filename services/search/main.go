package main

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/urfave/negroni"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/services"
)

func main() {

	//Register Logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	//config.ReadStandardConfig("search", logger)
	config.ReadStandardConfig("search", logger)
	serviceconfig.ConfigureDefaults()

	//Gorilla Router
	r := mux.NewRouter()

	//Bind to environment variables
	viper.SetEnvPrefix("search")
	_ = viper.BindEnv("elasticsearch_ca")

	//elasticSearchClient, err := elasticsearch.NewDefaultClient()
	elasticSearchClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:         []string{viper.GetString(serviceconfig.ElasticSearchAddress)},
		Username:          viper.GetString(serviceconfig.ElasticSearchUsername),
		Password:          viper.GetString(serviceconfig.ElasticSearchPassword),
		EnableDebugLogger: !viper.GetBool(config.Prod),
		CACert:            []byte(viper.GetString("elasticsearch_ca")),
	})

	if err != nil {
		logger.Fatalw("Elasticsearch Connection Error", "error", err)
	}

	gameSearch := services.NewGameSearch(logger, elasticSearchClient)
	gameSearchHandler := handlers.NewGameHandler(logger, gameSearch)

	//Register routes
	r.HandleFunc("/api/search", gameSearchHandler.GameSearch).Methods(http.MethodGet).Queries("q", "{.}")

	logger.Debugf("Routes registered successfully")

	//Register Middleware
	n := negroni.New(mw.NewLogger(logger.Desugar()), negroni.NewRecovery())
	n.UseHandler(r)

	//Configure HTTP server
	port := viper.GetString(config.HTTPPort)
	logger.Warnf("HTTP Port set to %v", port)
	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: viper.GetDuration(config.WriteTimeout),
		ReadTimeout:  viper.GetDuration(config.ReadTimeout),
		IdleTimeout:  viper.GetDuration(config.IdleTimeout),
		Handler:      n,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}
