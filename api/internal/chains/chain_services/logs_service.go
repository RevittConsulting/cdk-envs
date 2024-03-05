package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"log"
	"time"
)

type LogsService struct {
	Config    *config.Chains
	RpcConfig *config.RPCConfig
	Ticker    *time.Ticker

	MostRecentL1Block uint64
}

func NewLogsService(Config *config.Chains, RpcConfig *config.RPCConfig) *LogsService {
	ticker := time.NewTicker(5 * time.Second)
	return &LogsService{
		Config:    Config,
		RpcConfig: RpcConfig,
		Ticker:    ticker,
	}
}

func (s *LogsService) Start() error {
	clientL1 := jsonrpc.NewClient(s.RpcConfig.Url)

	log.Println("logs service started")
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
	log.Println("logs service stopped")
	s.Ticker.Stop()
	return nil
}

func (s *LogsService) GetMostRecentL1Block() uint64 {
	return s.MostRecentL1Block
}

func (s *LogsService) filterLogs(clientL1 *jsonrpc.Client, blockNum uint64) error {
	fromBlock := fmt.Sprintf("0x%X", blockNum-100)
	toBlock := "latest"
	address := interface{}(s.Config.Chains[ActiveChainConfigName].RollupAddress)
	topics := []interface{}{
		s.Config.Chains[ActiveChainConfigName].TopicsVerification,
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
