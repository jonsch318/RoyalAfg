package utils

import (
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// StartGrpcGracefully starts a grpc server on another go routine so when an interrupt signal hits,
// a timout is initialized and all active requests can be handled before shutting down.
func StartGrpcGracefully(logger *zap.SugaredLogger, server *grpc.Server, listener net.Listener) {
	go func() {
		if err := server.Serve(listener); err != nil {
			logger.Fatalw("Grpc server counld not listen", "error", err)
		}
	}()

	logger.Warnf("Grpc server Listening on address %v", listener.Addr())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	server.GracefulStop()

	logger.Warn("Http server shut down")
}
