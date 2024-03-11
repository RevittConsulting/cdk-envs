package chain_services

import (
	"fmt"
	"testing"
	"time"
)

type TestService struct {
	Ticker   *time.Ticker
	stopChan chan struct{}
}

func NewTestService() *TestService {
	ticker := time.NewTicker(5 * time.Second)
	return &TestService{
		Ticker:   ticker,
		stopChan: make(chan struct{}),
	}
}

func (s *TestService) Start() error {
	go func() {
		for {
			select {
			case <-s.Ticker.C:
				fmt.Println("test service tick")
			case <-s.stopChan:
				return
			}
		}
	}()

	return nil
}

func (s *TestService) Stop() error {
	close(s.stopChan)
	s.Ticker.Stop()
	return nil
}

func TestRuntime(t *testing.T) {
	registry := NewRegistry()
	testService := NewTestService()
	registry.Register("test", testService)
	r := NewRuntime(registry)

	err := r.StartServices("test")
	if err != nil {
		t.Errorf("error starting services: %v", err)
	}

	err = r.StopServices()
	if err != nil {
		t.Errorf("error stopping services: %v", err)
	}
}
