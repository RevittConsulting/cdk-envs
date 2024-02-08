package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type jsonRPCRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

func GetMostRecentBlock() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	client := &http.Client{}
	for {
		select {
		case <-ticker.C:
			blockNumber, err := getBlockNumber(client)
			if err != nil {
				fmt.Printf("Error getting block number: %v\n", err)
				return
			}

			fmt.Println(blockNumber)
		}
	}
}

func getBlockNumber(client *http.Client) (int, error) {
	requestBody, err := json.Marshal(jsonRPCRequest{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      1,
	})
	if err != nil {
		return 0, err
	}

	resp, err := client.Post("https://sepolia.publicgoods.network", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var response struct {
		Result string `json:"result"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	blockNumber, err := strconv.ParseInt(response.Result, 0, 64)
	if err != nil {
		return 0, err
	}

	return int(blockNumber), nil
}
