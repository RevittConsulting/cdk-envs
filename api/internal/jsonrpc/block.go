package jsonrpc

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

type Block struct {
	Number int64 `json:"number"`
}

func (c *RPCClient) GetMostRecentBlock() {
	ticker := time.NewTicker(c.cfg.PollingInterval)
	defer ticker.Stop()

	for range ticker.C {
		blockNumber, err := getBlockNumber(c)
		if err != nil {
			fmt.Printf("Error getting block number: %v\n", err)
			continue
		}

		block, err := saveBlockToDb(c.db, blockNumber)
		if err != nil {
			fmt.Printf("Error saving block to db: %v\n", err)
			continue
		}

		fmt.Println(block)
	}
}

func getBlockNumber(c *RPCClient) (int64, error) {
	requestBody := &jsonRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      1,
	}

	res, err := c.DoRequest(requestBody)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var response jsonRPCResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return 0, err
	}

	blockNumber := new(big.Int)
	blockNumber, ok := blockNumber.SetString(response.Result.(string), 0)
	if !ok {
		return 0, fmt.Errorf("Error converting block number to big.Int")
	}

	return blockNumber.Int64(), nil
}

func saveBlockToDb(db *bolt.DB, blockNumber int64) (*Block, error) {
	block := &Block{Number: blockNumber}

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Blocks"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(block)
		if err != nil {
			return err
		}

		key := fmt.Sprintf("%d", blockNumber) // Convert block number to string for key
		return b.Put([]byte(key), encoded)
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}
