package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/services"
)

func main() {

	//logging
	logger := log.RegisterService()
	defer log.CleanLogger()

	//configuration
	//config.ReadStandardConfig("search", logger)

	//Gorilla Routing
	r := mux.NewRouter()

	elasticSearchClient, err := elasticsearch.NewDefaultClient()

	if err != nil {
		logger.Fatalw("Elasticsearch Connection Error", "error", err)
	}

	gameSearch := services.NewGameSearch(logger, elasticSearchClient)
	gameSearchHandler := handlers.NewGameHandler(logger, gameSearch)

	gr := r.Methods(http.MethodGet).Subrouter()

	loggerHandler := mw.NewLoggerHandler(logger)
	stdChain := alice.New(loggerHandler.LogRoute)

	gr.Path("/api/search").Queries("q", "{.}").Handler(stdChain.ThenFunc(gameSearchHandler.GameSearch))

	// Start Application
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}
