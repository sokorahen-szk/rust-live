package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/sokorahen-szk/rust-live/api/proto"
	cfg "github.com/sokorahen-szk/rust-live/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	controller "github.com/sokorahen-szk/rust-live/internal/adapter/controller"
	"github.com/sokorahen-szk/rust-live/pkg/logger"

	"github.com/go-co-op/gocron"
	applicationBatch "github.com/sokorahen-szk/rust-live/internal/application/batch"
)

func main() {
	c := cfg.NewConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		logger.Fatalf("failed server binding port %d", c.Port)
	}

	scheduler(context.Background(), c, time.Local)

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterLiveServiceServer(server, &controller.LiveController{})

	logger.Infof("server starting port: %d", c.Port)
	if err := server.Serve(listener); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func scheduler(ctx context.Context, cfg *cfg.Config, tmLocation *time.Location) {
	if cfg.IsTest() {
		return
	}

	s := gocron.NewScheduler(tmLocation)
	s.Every(1).Minutes().Do(func(ctx context.Context) error {
		twitchFetchLiveVideosUsecase := applicationBatch.NewInjectTwitchFetchLiveVideosUsecase(ctx)
		err := twitchFetchLiveVideosUsecase.Handle(ctx)
		if err != nil {
			return err
		}

		updateLiveVideosUsecase := applicationBatch.NewInjectTwitchUpdateLiveVideosUsecase(ctx)
		err = updateLiveVideosUsecase.Handle(ctx)
		if err != nil {
			return err
		}

		return nil
	}, ctx)

	s.StartAsync()
}
