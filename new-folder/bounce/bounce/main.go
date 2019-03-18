package main

import (
	"net"
	"relay/reports/bounce/bounce/server"
	"relay/reports/bounce/proto/bounce"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.WithError(err).Error("could not start service")
		return
	}
	log.Info("starting service on :8080")

	srv := grpc.NewServer()
	log.Info("Bounce server is regitered here as well")
	bounce.RegisterBounceServer(srv, &server.Server{})
	srv.Serve(listener)
}
