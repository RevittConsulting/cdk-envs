package jsonrpc

import (
	"encoding/json"
	"math/big"
)

func (c *Client) ZkGetBatchNumber() (uint64, error) {
	res, err := c.DoRequest("zkevm_batchNumber")
	if err != nil {
		return 0, err
	}

	var result string
	err = json.Unmarshal(res.Result, &result)
	if err != nil {
		return 0, err
	}

	blockNum, ok := new(big.Int).SetString(result[2:], 16)
	if !ok {
		return 0, err
	}

	return blockNum.Uint64(), nil
}
