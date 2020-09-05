package docs

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/docs/pkg/docs/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Start starts the documentation service to serve the swagger documentation
func Start() {

	config.ConfigureDefaults()

	logger := log.NewLogger()
	logger.Warn("User service now running")

	defer logger.Warn("User service shut down")
	defer logger.Desugar().Sync()

	r := mux.NewRouter()

	gr := r.Methods(http.MethodGet).Subrouter()

	opts := middleware.RedocOpts{BasePath: "/docs", Path: "/", SpecURL: viper.GetString("SwaggerDocs.SwaggerUrl"), Title: viper.GetString("SwaggerDocs.Title")}
	//opts := middleware.RedocOpts{BasePath: "/", Path: "/", SpecURL: "http://localhost:9000/docs/swagger.yaml", Title: viper.GetString("SwaggerDocs.Title")}
	gr.Handle("/docs", middleware.Redoc(opts, nil))
	//gr.Handle("/docs/swagger.yaml", http.FileServer(http.Dir("./")))
	gr.HandleFunc("/docs/swagger.yaml", func(rw http.ResponseWriter, r *http.Request) {
		logger.Infof("Path /docs/swagger.yaml called")
		http.ServeFile(rw, r, "./swagger.yaml")
	})

	server := &http.Server{
		Addr:         ":" + viper.GetString(config.Port),
		WriteTimeout: viper.GetDuration(config.WriteTimeout),
		ReadTimeout:  viper.GetDuration(config.ReadTimeout),
		IdleTimeout:  viper.GetDuration(config.IdleTimeout),
		Handler:      r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulTimeout))
}
