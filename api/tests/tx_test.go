package tests

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/internal/tx"
	"testing"
)

func TestDoTx(t *testing.T) {
	t.Skip("Skipping test")

	httpService := tx.NewService(nil)

	req := &tx.Request{
		Host:        "http://localhost:8545",
		ToAddress:   "",
		FromAddress: "",
		PrivateKey:  "",
		Amount:      10000,
		GasLimit:    21000,
		GasPrice:    1e9,
	}

	output, err := httpService.CreateTx(context.Background(), req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if output == nil {
		t.Errorf("Output is nil")
	}
}
