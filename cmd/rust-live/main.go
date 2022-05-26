package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-co-op/gocron"
	pb "github.com/sokorahen-szk/rust-live/api/proto"
	cfg "github.com/sokorahen-szk/rust-live/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	controller "github.com/sokorahen-szk/rust-live/internal/adapter/controller"
	batch "github.com/sokorahen-szk/rust-live/internal/domain/batch/application"
	"github.com/sokorahen-szk/rust-live/pkg/logger"
)

func main() {
	c := cfg.NewConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		logger.Fatalf("failed server binding port %d", c.Port)
	}

	scheduler()

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterLiveServiceServer(server, &controller.LiveController{})

	logger.Infof("server starting port: %d", c.Port)
	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func scheduler() {
	ctx := context.Background()
	s := gocron.NewScheduler(time.Local)

	s.Every(1).Minutes().Do(func(ctx context.Context) error {
		usecase := batch.NewInjectFetchLiveVideosUsecase(ctx)
		return usecase.Handle(ctx)
	}, ctx)

	s.Every(1).Minutes().Do(func(ctx context.Context) error {
		usecase := batch.NewInjectUpdateLiveVideosUsecase(ctx)
		return usecase.Handle(ctx)
	}, ctx)

	s.StartAsync()
}
