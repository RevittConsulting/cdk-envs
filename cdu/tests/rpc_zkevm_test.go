package tests

import (
	"github.com/RevittConsulting/chain-dev-utils/internal/jsonrpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BatchNumber(t *testing.T) {
	client := jsonrpc.NewClient("https://rpc.cardona.zkevm-rpc.com/")

	batchNum, err := client.ZkGetBatchNumber()
	if err != nil {
		t.Error(err)
	}

	t.Log("batchNum:", batchNum)

	assert.Equal(t, batchNum > 0, true)
}
