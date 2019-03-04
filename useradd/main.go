package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/proto/useradd"
	"github.com/wizelineacademy/GoWorkshop/useradd/server"
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
	useradd.RegisterUserAddServer(srv, &server.Server{})
	srv.Serve(listener)
}
