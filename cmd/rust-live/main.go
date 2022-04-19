package main

import (
	"fmt"
	"net"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sokorahen-szk/rust-live/internal/application"
	log "github.com/sokorahen-szk/rust-live/pkg/logger"
)

const server_port = 9000

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port))
	if err != nil {
		log.Fatalf("failed server binding port %d", server_port)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterLiveServiceServer(server, &application.LiveService{})

	log.Infof("server starting port: %d", server_port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
