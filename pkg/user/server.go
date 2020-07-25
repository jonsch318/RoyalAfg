package user

import (
	"github.com/JohnnyS318/RoyalAfgInGo/internal/log"
	"github.com/gorilla/mux"
	"net/http"
)

func Start(){
	logger := log.NewLogger()

	logger.Warn("Application started. Router will be configured next")

	r := mux.NewRouter()
	r.HandleFunc("/", func (response http.ResponseWriter, request *http.Request){
		logger.Info("Root called")
	}).Methods(http.MethodGet)
}
