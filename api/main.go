package main

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	app.Start(start)
}

func start(ctx context.Context, r *chi.Mux) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	log.Fatal(http.ListenAndServe(cfg.Port, r))

	return nil
}
