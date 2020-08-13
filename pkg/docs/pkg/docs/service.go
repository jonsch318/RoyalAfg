package docs

import (
	"fmt"
	"net/http"

	"royalafg/pkg/docs/pkg/docs/config"
	"royalafg/pkg/shared/pkg/log"
	"royalafg/pkg/shared/pkg/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Start starts the documentation service to serve the swagger documentation
func Start() {

	logger := log.NewLogger()
	logger.Warn("User service now running")

	defer logger.Warn("User service shut down")
	defer logger.Desugar().Sync()

	r := mux.NewRouter()

	gr := r.Methods(http.MethodGet).Subrouter()

	port := viper.GetString(config.Port)

	opts := middleware.RedocOpts{SpecURL: fmt.Sprintf("http://localhost:%v/swagger.yaml", port), Title: viper.GetString("SwaggerDocs.Title")}
	gr.Handle("/docs", middleware.Redoc(opts, nil))
	gr.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         ":" + viper.GetString(config.Port),
		WriteTimeout: viper.GetDuration(config.WriteTimeout),
		ReadTimeout:  viper.GetDuration(config.ReadTimeout),
		IdleTimeout:  viper.GetDuration(config.IdleTimeout),
		Handler:      r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulTimeout))
}
