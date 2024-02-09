package jsonrpc

import (
	"github.com/RevittConsulting/cdk-envs/config"
	"testing"
)

func Test_GetLogs(t *testing.T) {
	cfg := &config.RPCConfig{
		Url2: "https://rpc2.sepolia.org",
	}
	client := NewRPCClient(cfg, nil)

	highestBlock, err := getBlockNumber(client)
	if err != nil {
		t.Errorf("Error getting block number: %v", err)
	}

	batch, err := getLogs(client, highestBlock)
	if err != nil {
		t.Errorf("Error getting logs: %v", err)
	}

	t.Log(batch)
}
