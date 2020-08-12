package utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

func StartGracefully(logger *zap.SugaredLogger, server *http.Server, timeoutDuration time.Duration) {
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalw("Http server counld not listen and serve", "error", err)
		}
	}()

	logger.Warnf("Http server Listening on address %v", server.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	err := server.Shutdown(ctx)

	if err != nil {
		logger.Fatalw("Http server encouterd an error while shutting down", "error", err)
	}

	logger.Warn("Http server shutting down")
}

// ToJSON encodes objects and writes them using the specified io.Writer
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

// FromJSON decodes objects from an io.Reader
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}

// FromFormURLEncoded decodes objects form a parsed form
func FromFormURLEncoded(i interface{}, r map[string][]string) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(i, r)
}

// FromFormURLEncodedRequest decodes objects from a form url encoded request
func FromFormURLEncodedRequest(i interface{}, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return FromFormURLEncoded(i, r.Form)
}
