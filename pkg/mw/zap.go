package mw

import (
	"net/http"
	"time"

	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

//Logger is a Middleware for logging using the zap logger
type Logger struct {
	zap *zap.Logger
}

//NewLogger creates a new zap logging middleware
func NewLogger(zap *zap.Logger) *Logger {
	return &Logger{
		zap: zap,
	}
}

//ServeHTTP is the middleware handler for the logging middleware.
func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)

	if res.Status() < 400 {
		l.zap.Info("MW Log",
			zap.Time("StartTime", start),
			zap.Int("Status", res.Status()),
			zap.Duration("Duration", time.Since(start)),
			zap.String("Host", r.Host),
			zap.String("Method", r.Method),
			zap.String("Path", r.URL.Path),
		)
	} else {
		l.zap.Error("MW Log Error",
			zap.Time("StartTime", start),
			zap.Int("Status", res.Status()),
			zap.Duration("Duration", time.Since(start)),
			zap.String("Host", r.Host),
			zap.String("Method", r.Method),
			zap.String("Path", r.URL.Path),
		)
	}
}