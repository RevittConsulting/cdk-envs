package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"log"
	"strconv"
	"strings"
	"time"
)

type LogsService struct {
	Config      *config.Chains
	L1Contracts *config.L1Contracts
	RpcConfig   *config.RPCConfig
	Ticker      *time.Ticker

	stopChan chan struct{}

	MostRecentL1Block     uint64
	HighestSequencedBatch uint64
}

func NewLogsService(Config *config.Chains, L1Contracts *config.L1Contracts, RpcConfig *config.RPCConfig) *LogsService {
	ticker := time.NewTicker(5 * time.Second)
	return &LogsService{
		Config:      Config,
		L1Contracts: L1Contracts,
		RpcConfig:   RpcConfig,
		Ticker:      ticker,
		stopChan:    make(chan struct{}),
	}
}

func (s *LogsService) Start() error {
	s.Ticker = time.NewTicker(5 * time.Second)
	s.stopChan = make(chan struct{})

	clientL1 := jsonrpc.NewClient(s.RpcConfig.Url)

	go func() {
		log.Println("logs service started")
		defer log.Println("logs service stopped")
		for {
			select {
			case <-s.Ticker.C:
				blockNum, err := clientL1.EthBlockNumber()
				if err != nil {
					fmt.Println("error getting most recent block")
				}
				if blockNum > s.MostRecentL1Block {
					err = s.filterLogs(clientL1, blockNum)
					if err != nil {
						fmt.Println("error filtering logs")
					}
					s.MostRecentL1Block = blockNum
				}
			case <-s.stopChan:
				return
			}
		}
	}()

	return nil
}

func (s *LogsService) Stop() error {

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

func (s *LogsService) GetHighestSequencedBatch() uint64 {
	return s.HighestSequencedBatch
}

func (s *LogsService) filterLogs(clientL1 *jsonrpc.Client, blockNum uint64) error {
	fromBlock := fmt.Sprintf("0x%X", blockNum-100)
	toBlock := "latest"
	address := interface{}(s.Config.Chains[ActiveChainConfigName].RollupAddress)
	topics := []interface{}{
		s.Config.Chains[ActiveChainConfigName].TopicsSequence,
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

	if len(logs) == 0 {
		return nil
	}

	sequencedBatch, err := s.getHighestSequencedBatch(logs[0].Topics[1])
	if err != nil {
		return fmt.Errorf("error getting highest sequenced batch: %w", err)
	}

	if sequencedBatch > s.HighestSequencedBatch {
		s.HighestSequencedBatch = sequencedBatch
	}

	return nil
}

func (s *LogsService) getHighestSequencedBatch(topic string) (uint64, error) {
	batchNum, err := hashToUint64(topic)
	if err != nil {
		return 0, fmt.Errorf("error converting hash to uint64: %w", err)
	}
	return batchNum, nil
}

func hashToUint64(hash string) (uint64, error) {
	cleanHash := strings.TrimPrefix(hash, "0x")
	return strconv.ParseUint(cleanHash, 16, 64)
}
