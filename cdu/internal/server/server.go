package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/RevittConsulting/chain-dev-utils/config"
	"github.com/RevittConsulting/chain-dev-utils/internal/chains/chain_services"
	"github.com/RevittConsulting/chain-dev-utils/pkg/atomics"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Config        *config.Config
	ShuttingDown  *atomics.AtomicBool
	Router        *chi.Mux
	Deps          *dependencies
	Signal        chan os.Signal
	ChainServices *chain_services.Registry
}

func (s *Server) Init() error {
	var shutdown atomics.AtomicBool
	shutdown.Set(false)
	s.ShuttingDown = &shutdown

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	s.Signal = c

	return nil
}

func (s *Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := s.Setup(); err != nil {
		return err
	}

	port := fmt.Sprintf(":%v", s.Config.Port)
	server := http.Server{
		Addr:    port,
		Handler: s.Router,
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			if err = s.ChainServices.StopAll(); err != nil {
				return err
			}
			return err
		}
		return nil
	})

	g.Go(func() error {
		select {
		case <-s.Signal:
			s.ShuttingDown.Set(true)
			time.Sleep(time.Duration(s.Config.ShutdownTime) * time.Second)
			cancel()
			return server.Shutdown(context.Background())
		case <-ctx.Done():
			s.ShuttingDown.Set(true)
			return server.Shutdown(context.Background())
		}
	})

	fmt.Println("server started on port " + port)

	if err := g.Wait(); err != nil {
		return err
	}

	if err := s.ChainServices.StopAll(); err != nil {
		return err
	}

	fmt.Println("server gracefully stopped")
	return nil
}

func (s *Server) Setup() error {
	if err := s.SetupDeps(); err != nil {
		return err
	}

	if err := s.SetupHandlers(); err != nil {
		return err
	}

	return nil
}
