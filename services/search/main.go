package main

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	serviceconfig "github.com/JohnnyS318/RoyalAfgInGo/services/search/config"
	"github.com/gorilla/mux"
	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/viper"
)

func main() {

	//config
	logger := log.RegisterService()
	defer log.CleanLogger()

	config.ReadStandardConfig("search", logger)

	viper.SetEnvPrefix(serviceconfig.ENV_PREFIX)

	r := mux.NewRouter()

	//register meilisearch

	meileClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   viper.GetString(serviceconfig.MeilisearchHost),
		APIKey: viper.GetString(serviceconfig.MeilisearchAPIKey),
	})

	meileClient.Index("").Search("test", &meilisearch.SearchRequest{})

}
