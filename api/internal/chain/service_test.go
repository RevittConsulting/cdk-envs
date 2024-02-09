package chain

import (
	"context"
	"testing"
)

func Test_GetChain(t *testing.T) {
	s := NewService(nil, nil)
	chains, err := s.GetChains(context.TODO())
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	if chains == nil {
		t.Errorf("Chains should not be nil")
	}
}
