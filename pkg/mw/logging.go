package mw

import (
	"fmt"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"

	"go.uber.org/zap"
)

// LoggerHandler Handler only for logger middleware
type LoggerHandler struct {
	l *zap.SugaredLogger
}

// NewLoggerHandler creates a new LoggerHandler
func NewLoggerHandler(logger *zap.SugaredLogger) *LoggerHandler {
	return &LoggerHandler{l: logger}
}

// LogRoute logs the called route together
func (h *LoggerHandler) LogRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		h.l.Debugw(fmt.Sprintf("%v called", r.URL.Path), "route", r.URL.Path)
		next.ServeHTTP(rw, r)
	})
}

// LogRouteWithIP logs the called route together with the ip
func (h *LoggerHandler) LogRouteWithIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		h.l.Debugw(fmt.Sprintf("%v called", r.URL.Path), "route", r.URL.Path, "remote_ip", utils.GetIP(r))
		next.ServeHTTP(rw, r)
	})
}

// ContentTypeJSON sets the content type and x-content-type-options to json
func (h *LoggerHandler) ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(rw, r)
	})

}
