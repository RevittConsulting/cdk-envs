package server

import (
	"github.com/RevittConsulting/cdk-envs/internal/chain"
	"github.com/RevittConsulting/cdk-envs/internal/health"
	"github.com/boltdb/bolt"
	"github.com/go-chi/chi/v5"
)

type dependencies struct {
	chain *chain.Service
}

func (s *Server) SetupDeps(db *bolt.DB) error {
	var deps dependencies
	chainDb := chain.NewDb(db)
	deps.chain = chain.NewService(chainDb, &s.Config.Chain)
	s.Deps = &deps
	return nil
}

func (s *Server) SetupHandlers() error {
	s.Router.Route("/api/v1", func(r chi.Router) {
		health.NewHandler(r, s.ShuttingDown)
		chain.NewHandler(r, s.Deps.chain)
	})
	return nil
}
