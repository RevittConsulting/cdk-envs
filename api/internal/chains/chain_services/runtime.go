package chain_services

import (
	"errors"
	"log"
	"sync"
)

var (
	ErrServiceAlreadyRunning = errors.New("a service is already running")
	ErrServiceNotFound       = errors.New("service not found")
)

var (
	ActiveChainConfigName string
)

type Runtime struct {
	ChainServices *Registry

	running   bool
	mu        sync.Mutex
	waitGroup sync.WaitGroup

	ActiveServices []IService
}

func NewRuntime(ChainServices *Registry) *Runtime {
	return &Runtime{
		ChainServices: ChainServices,
		running:       false,
	}
}

func (r *Runtime) StartServices(serviceName ...string) error {
	r.mu.Lock()
	if r.running {
		r.mu.Unlock()
		return ErrServiceAlreadyRunning
	}
	r.mu.Unlock()

	for _, service := range serviceName {
		s, err := r.ChainServices.GetService(service)
		if err != nil {
			return err
		}

		if err = s.Start(); err != nil {
			return err
		}

		r.ActiveServices = append(r.ActiveServices, s)
	}

	r.running = true

	return nil
}

func (r *Runtime) StopServices() error {
	r.mu.Lock()
	if !r.running {
		r.mu.Unlock()
		return nil
	}
	r.mu.Unlock()

	services := r.ActiveServices
	for _, service := range services {
		if err := service.Stop(); err != nil {
			return err
		}
	}
	r.ActiveServices = nil

	r.running = false

	return nil
}

func (r *Runtime) RestartService(chainName string) error {
	ActiveChainConfigName = chainName
	log.Println("restarting service")

	if r.ActiveServices != nil {
		for _, service := range r.ActiveServices {
			if err := service.Stop(); err != nil {
				return err
			}
		}
	}

	return r.StartServices(Logs)
}

func (r *Runtime) GetActiveServices() []IService {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.ActiveServices
}
