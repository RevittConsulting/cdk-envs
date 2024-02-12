package jsonrpc

import (
	"encoding/json"
	"math/big"
)

func (c *Client) EthBlockNumber() (uint64, error) {
	res, err := c.DoRequest("eth_blockNumber")
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

func (c *Client) EthGetLogs(query LogQuery) ([]Log, error) {
	res, err := c.DoRequest("eth_getLogs", query)
	if err != nil {
		return nil, err
	}

	var logs []Log
	err = json.Unmarshal(res.Result, &logs)
	if err != nil {
		return nil, err
	}

	return logs, nil
}
