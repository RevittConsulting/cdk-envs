package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/chain-dev-utils/config"
	"github.com/RevittConsulting/chain-dev-utils/internal/jsonrpc"
	"github.com/RevittConsulting/chain-dev-utils/pkg/hexadecimal"
	"log"
	"math/big"
	"time"
)

type LogsService struct {
	Config      *config.Chains
	L1Contracts *config.L1Contracts
	RpcConfig   *jsonrpc.Config
	Ticker      *time.Ticker

	stopChan chan struct{}

	MostRecentL1Block     uint64
	HighestSequencedBatch uint64
	HighestVerifiedBatch  uint64
}

func NewLogsService(Config *config.Chains, L1Contracts *config.L1Contracts, RpcConfig *jsonrpc.Config) *LogsService {
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

	clientL1 := jsonrpc.NewClient(s.Config.Chains[ActiveChainConfigName].L1RpcUrl)

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
					err = s.filterLogsSequence(clientL1, blockNum)
					if err != nil {
						fmt.Println("error filtering logs")
					}
					err = s.filterLogsVerification(clientL1, blockNum)
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

func (s *LogsService) GetHighestVerifiedBatch() uint64 {
	return s.HighestVerifiedBatch
}

func (s *LogsService) filterLogsSequence(clientL1 *jsonrpc.Client, blockNum uint64) error {
	latestBlock := big.NewInt(int64(blockNum))
	blockRange := big.NewInt(10000)
	var foundLog *jsonrpc.Log

	for {
		from := new(big.Int).Sub(latestBlock, blockRange)
		if from.Cmp(big.NewInt(0)) == -1 {
			from = big.NewInt(0)
		}

		fromBlock := fmt.Sprintf("0x%X", from)
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

		if len(logs) > 0 {
			fmt.Println("Found logs.")
			foundLog = &logs[len(logs)-1]
			break
		}

		latestBlock = new(big.Int).Sub(from, big.NewInt(1))
		if latestBlock.Cmp(big.NewInt(0)) == -1 || from.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("No logs found.")
			return nil
		}

		fmt.Println("No logs found. Continuing to search.")
	}

	var sequencedBatch uint64
	if foundLog != nil {
		fmt.Printf("Found log: %+v\n", *foundLog)
		if len(foundLog.Topics) >= 2 {
			sequenced, err := hexadecimal.HashToUint64(foundLog.Topics[1])
			if err != nil {
				return fmt.Errorf("error getting highest sequenced batch: %w", err)
			}
			sequencedBatch = sequenced
			fmt.Println("The found log does not contain a second topic.")
		}
	} else {
		fmt.Println("No logs found.")
	}

	if sequencedBatch > s.HighestSequencedBatch {
		s.HighestSequencedBatch = sequencedBatch
	}

	return nil
}

func (s *LogsService) filterLogsVerification(clientL1 *jsonrpc.Client, blockNum uint64) error {
	latestBlock := big.NewInt(int64(blockNum))
	blockRange := big.NewInt(10000)
	var foundLog *jsonrpc.Log

	for {
		from := new(big.Int).Sub(latestBlock, blockRange)
		if from.Cmp(big.NewInt(0)) == -1 {
			from = big.NewInt(0)
		}

		fromBlock := fmt.Sprintf("0x%X", from)
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

		if len(logs) > 0 {
			fmt.Println("Found logs.")
			foundLog = &logs[len(logs)-1]
			break
		}

		latestBlock = new(big.Int).Sub(from, big.NewInt(1))
		if latestBlock.Cmp(big.NewInt(0)) == -1 || from.Cmp(big.NewInt(0)) == 0 {
			fmt.Println("No logs found.")
			return nil
		}

		fmt.Println("No logs found. Continuing to search.")
	}

	var verifiedBatch uint64
	if foundLog != nil {
		fmt.Printf("Found log: %+v\n", *foundLog)
		if len(foundLog.Topics) >= 2 {
			verify, err := hexadecimal.HashToUint64(foundLog.Topics[1])
			if err != nil {
				return fmt.Errorf("error getting highest sequenced batch: %w", err)
			}
			verifiedBatch = verify
		} else {
			fmt.Println("The found log does not contain a second topic.")
		}
	} else {
		fmt.Println("No logs found.")
	}

	if verifiedBatch > s.HighestVerifiedBatch {
		s.HighestVerifiedBatch = verifiedBatch
	}

	return nil
}
