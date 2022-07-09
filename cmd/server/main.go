package main

import (
	"net"

	"github.com/shani1998/shani1998.github.io/pkg/pb"
	"github.com/shani1998/shani1998.github.io/pkg/svc"
	log "github.com/sirupsen/logrus"
)

const (
	Addr = "127.0.0.1:50051"
)

func main() {
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Infof("Listening on %s", Addr)
	server := NewGRPCServer()
	pb.RegisterPortfolioServer(server, &svc.PortfolioServer{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
