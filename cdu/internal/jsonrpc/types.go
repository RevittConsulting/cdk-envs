package jsonrpc

import (
	"encoding/json"
	"strconv"
)

type Request struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	JsonRpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
}

type LogQuery struct {
	BlockHash *string        `json:"blockHash,omitempty"`
	FromBlock *string        `json:"fromBlock,omitempty"`
	ToBlock   *string        `json:"toBlock,omitempty"`
	Address   *interface{}   `json:"address,omitempty"`
	Topics    *[]interface{} `json:"topics,omitempty"`
}

type Log struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber HexUint  `json:"blockNumber"`
	TxHash      string   `json:"transactionHash"`
	TxIndex     HexUint  `json:"transactionIndex"`
	BlockHash   string   `json:"blockHash"`
	LogIndex    HexUint  `json:"logIndex"`
	Removed     bool     `json:"removed"`
}

type HexUint uint

func (h *HexUint) UnmarshalJSON(data []byte) error {
	str := string(data[1 : len(data)-1])

	val, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		return err
	}

	*h = HexUint(val)
	return nil
}

type Block struct {
	Number int64 `json:"number"`
}
