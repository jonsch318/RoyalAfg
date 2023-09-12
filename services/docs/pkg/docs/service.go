package docs

import (
	"net/http"

	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/pkg/utils"
	"github.com/jonsch318/royalafg/services/docs/pkg/docs/serviceconfig"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Start starts the documentation service to serve the swagger documentation
func Start() {
	logger := log.RegisterService()
	defer log.CleanLogger()
	config.ReadStandardConfig("docs", logger)
	serviceconfig.ConfigureDefaults()

	r := mux.NewRouter()

	gr := r.Methods(http.MethodGet).Subrouter()

	opts := middleware.RedocOpts{BasePath: "/docs", Path: "/", SpecURL: viper.GetString(serviceconfig.SwaggerUrl), Title: viper.GetString(serviceconfig.SwaggerTitle)}
	gr.Handle("/docs", middleware.Redoc(opts, nil))
	gr.HandleFunc("/docs/swagger.yaml", func(rw http.ResponseWriter, r *http.Request) {
		logger.Infof("Path /docs/swagger.yaml called service")
		http.ServeFile(rw, r, viper.GetString(serviceconfig.SwaggerFile))
	})

	server := &http.Server{
		Addr:    ":" + viper.GetString(config.HTTPPort),
		Handler: r,
	}

	utils.StartGracefully(logger, server, viper.GetDuration(config.GracefulShutdownTimeout))
}
