package chain_services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestService struct{}

func (s *TestService) Start() error {
	return nil
}

func (s *TestService) Stop() error {
	return nil
}

func TestRuntime(t *testing.T) {
	registry := NewRegistry()
	registry.Register("test", &TestService{})
	r := NewRuntime(registry)

	// Stop service
	err := r.StopService()
	if err != nil {
		t.Errorf("Error stopping service: %v", err)
	}

	// Start service
	err = r.StartService("test")
	if err != nil {
		t.Errorf("Error starting service: %v", err)
	}

	// Start service again
	err = r.StartService("test")
	if !errors.Is(err, ErrServiceAlreadyRunning) {
		assert.Equal(t, ErrServiceAlreadyRunning, err, "starting same service again")
	}

	// Stop service
	err = r.StopService()
	if err != nil {
		t.Errorf("Error stopping service: %v", err)
	}

	// Start invalid service
	err = r.StartService("invalid-service")
	if !errors.Is(err, ErrServiceNotFound) {
		assert.Equal(t, ErrServiceNotFound, err, "starting invalid service")
	}
}
