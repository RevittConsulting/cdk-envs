package server

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/internal/health"
	"github.com/go-chi/chi/v5"
)

func (s *Server) SetupHandlers(ctx context.Context) error {
	s.Router.Route("/api/v1", func(r chi.Router) {
		health.NewHandler(r, s.ShuttingDown)
	})
	return nil
}
