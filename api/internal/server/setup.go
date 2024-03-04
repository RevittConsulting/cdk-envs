package server

import (
	"github.com/RevittConsulting/cdk-envs/internal/buckets"
	"github.com/RevittConsulting/cdk-envs/internal/buckets/db/mdbx"
	"github.com/RevittConsulting/cdk-envs/internal/chains"
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
	"github.com/RevittConsulting/cdk-envs/internal/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type dependencies struct {
	chain   *chains.Service
	buckets *buckets.Service
}

func (s *Server) SetupDeps() error {
	var deps dependencies

	// chain routine services
	registry := chain_services.NewRegistry()

	// block service (gets most recent block)
	blockService := chain_services.NewBlockService(&s.Config.RPC)
	registry.Register(chain_services.Block, blockService)

	// logs service
	logsService := chain_services.NewLogsService(&s.Config.RPC)
	registry.Register(chain_services.Logs, logsService)

	// chain http service
	deps.chain = chains.NewService(s.Config.Chains, registry)

	// datacryp
	mdbxdb := mdbx.New()
	deps.buckets = buckets.NewService(s.Config.Buckets, mdbxdb)

	s.Deps = &deps
	s.ChainServices = registry
	return nil
}

func (s *Server) SetupHandlers() error {

	s.Router = chi.NewRouter()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	s.Router.Use(crs.Handler)

	s.Router.Route("/api/v1", func(r chi.Router) {
		health.NewHandler(r, s.ShuttingDown)
		chains.NewHandler(r, s.Deps.chain)
		buckets.NewHandler(r, s.Deps.buckets)
	})
	return nil
}
