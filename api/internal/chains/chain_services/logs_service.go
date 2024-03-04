package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"time"
)

type LogsService struct {
	RpcConfig *config.RPCConfig
	Ticker    *time.Ticker

	MostRecentL1Block uint64
}

func NewLogsService(RpcConfig *config.RPCConfig) *LogsService {
	ticker := time.NewTicker(5 * time.Second)
	return &LogsService{
		RpcConfig: RpcConfig,
		Ticker:    ticker,
	}
}

func (s *LogsService) Start() error {
	clientL1 := jsonrpc.NewClient(s.RpcConfig.Url)

	for range s.Ticker.C {
		blockNum, err := clientL1.EthBlockNumber()
		if err != nil {
			return fmt.Errorf("error getting most recent block: %w", err)
		}

		s.MostRecentL1Block = blockNum

		err = s.filterLogs(clientL1, blockNum)
		if err != nil {
			return fmt.Errorf("error filtering logs: %w", err)
		}
	}

	return nil
}

func (s *LogsService) Stop() error {
	s.Ticker.Stop()
	return nil
}

func (s *LogsService) GetMostRecentL1Block() uint64 {
	return s.MostRecentL1Block
}

func (s *LogsService) filterLogs(clientL1 *jsonrpc.Client, blockNum uint64) error {
	fromBlock := fmt.Sprintf("0x%X", blockNum-100)
	toBlock := "latest"
	address := interface{}("0xA13Ddb14437A8F34897131367ad3ca78416d6bCa")
	topics := []interface{}{
		"0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966",
	}

	query := jsonrpc.LogQuery{
		FromBlock: &fromBlock,
		ToBlock:   &toBlock,
		Address:   &address,
		Topics:    &topics,
	}

	logs, err := clientL1.EthGetLogs(query)
	if err != nil {
		return fmt.Errorf("error getting logs: %w", err)
	}

	fmt.Println("logs:", logs)

	return nil
}
