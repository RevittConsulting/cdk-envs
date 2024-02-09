package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/boltdb/bolt"
	"io"
	"net/http"
)

type jsonRPCRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type jsonRPCResponse struct {
	Id      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

type RPCClient struct {
	cfg *config.RPCConfig
	db  *bolt.DB
}

func NewRPCClient(cfg *config.RPCConfig, db *bolt.DB) *RPCClient {
	return &RPCClient{
		cfg: cfg,
		db:  db,
	}
}

func (c *RPCClient) DoRequest(req *jsonRPCRequest) ([]byte, error) {
	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.cfg.Url2, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
