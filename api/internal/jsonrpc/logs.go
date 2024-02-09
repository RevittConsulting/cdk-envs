package jsonrpc

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type logQuery struct {
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Address   string   `json:"address"`
	Topics    []string `json:"topics"`
}

type LogEntry struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

type LogResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []LogEntry `json:"result"`
}

func (c *RPCClient) GetLogs() {
	ticker := time.NewTicker(c.cfg.PollingInterval)
	defer ticker.Stop()

	for range ticker.C {
		highestBlock, err := getBlockNumber(c)
		if err != nil {
			fmt.Printf("Error getting block number: %v\n", err)
			continue
		}

		batch, err := getLogs(c, highestBlock)
		if err != nil {
			fmt.Printf("Error getting block number: %v\n", err)
			continue
		}

		fmt.Println(batch)

	}
}

func getLogs(client *RPCClient, highestBlock int64) (*LogResponse, error) {
	query := logQuery{
		FromBlock: fmt.Sprintf("0x%X", highestBlock-1000),
		ToBlock:   "latest",
		Address:   "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff", // polygon rollup manager
		Topics:    []string{"0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3"},
	}

	req := &jsonRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_getLogs",
		Params:  []interface{}{query},
		Id:      1,
	}

	res, err := client.DoRequest(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var response *LogResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return response, nil
}
