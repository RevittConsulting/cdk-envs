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

	stopChan chan struct{}

	MostRecentL1Block uint64
}

func NewLogsService(Config *config.Chains, RpcConfig *config.RPCConfig) *LogsService {
	ticker := time.NewTicker(5 * time.Second)
	return &LogsService{
		Config:    Config,
		RpcConfig: RpcConfig,
		Ticker:    ticker,
		stopChan:  make(chan struct{}),
	}
}

func (s *LogsService) Start() error {
	s.Ticker = time.NewTicker(5 * time.Second)
	s.stopChan = make(chan struct{})

	clientL1 := jsonrpc.NewClient(s.RpcConfig.Url)

	log.Println("logs service started")
	go func() {
		for {
			select {
			case <-s.Ticker.C:
				blockNum, err := clientL1.EthBlockNumber()
				if err != nil {
					fmt.Println("error getting most recent block")
				}

				s.MostRecentL1Block = blockNum

				err = s.filterLogs(clientL1, blockNum)
				if err != nil {
					fmt.Println("error filtering logs")
				}
			case <-s.stopChan:
				return
			}
		}
	}()

	return nil
}

func (s *LogsService) Stop() error {
	log.Println("logs service stopped")

	if s.stopChan != nil {
		close(s.stopChan)
		s.stopChan = nil
	}

	if s.Ticker != nil {
		s.Ticker.Stop()
		s.Ticker = nil
	}

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
