package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
)

type App struct{}

type StartFunc func(ctx context.Context, r *chi.Mux) error

func Start(startFunc StartFunc) {
	ctx := context.Background()

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)

	err := startFunc(ctx, r)
	if err != nil {
		log.Fatal("failed to start app")
	}
}
