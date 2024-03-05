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

	serviceChan chan IService
	running     bool
	mu          sync.Mutex
	waitGroup   sync.WaitGroup

	ActiveService IService
}

func NewRuntime(ChainServices *Registry) *Runtime {
	return &Runtime{
		ChainServices: ChainServices,
		serviceChan:   make(chan IService),
		running:       false,
	}
}

func (r *Runtime) StartService(serviceName string) error {
	r.mu.Lock()
	if r.running {
		return ErrServiceAlreadyRunning
	}
	r.running = true
	defer r.mu.Unlock()

	service, err := r.ChainServices.GetService(serviceName)
	if err != nil {
		return ErrServiceNotFound
	}

	r.waitGroup.Add(1)
	go func() {
		defer r.waitGroup.Done()

		r.ActiveService = service
		if err = service.Start(); err != nil {
			r.serviceChan <- nil
			return
		}

		r.serviceChan <- service // assign the chan in the routine, so we can stop it later
	}()

	return nil
}

func (r *Runtime) StopService() error {
	r.mu.Lock()
	if !r.running {
		r.mu.Unlock()
		return nil
	}
	r.running = false
	defer r.mu.Unlock()

	r.ActiveService = nil

	service := <-r.serviceChan
	if err := service.Stop(); err != nil {
		return err
	}

	r.waitGroup.Wait()

	return nil
}

func (r *Runtime) RestartService(chainName string) error {
	ActiveChainConfigName = chainName
	log.Println("restarting service")

	if r.ActiveService != nil {
		if err := r.StopService(); err != nil {
			return err
		}
	}

	return r.StartService(Logs)
}

func (r *Runtime) GetActiveService() IService {
	return r.ActiveService
}
