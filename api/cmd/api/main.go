package main

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/server"
)

func main() {
	server.Start(start)
}

func start(ctx context.Context, s *server.Server) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	return s.Setup(ctx, cfg)
}
