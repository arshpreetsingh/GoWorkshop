package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/bounce/server"
	"github.com/wizelineacademy/GoWorkshop/proto/bounce"
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
	bounce.RegisterBounceServer(srv, &server.Server{})
	srv.Serve(listener)
}
