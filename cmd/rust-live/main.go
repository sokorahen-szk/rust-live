package main

import (
	"fmt"
	"net"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	cfg "github.com/sokorahen-szk/rust-live/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	controller "github.com/sokorahen-szk/rust-live/internal/adapter/controller"
	log "github.com/sokorahen-szk/rust-live/pkg/logger"
)

func main() {
	c := cfg.NewConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Fatalf("failed server binding port %d", c.Port)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterLiveServiceServer(server, &controller.LiveController{})

	log.Infof("server starting port: %d", c.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
