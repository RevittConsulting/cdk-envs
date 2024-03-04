package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"time"
)

type BlockService struct {
	RpcConfig *config.RPCConfig
	Ticker    *time.Ticker

	MostRecentL1Block uint64
}

func NewBlockService(RpcConfig *config.RPCConfig) *BlockService {
	ticker := time.NewTicker(10 * time.Second)
	return &BlockService{
		RpcConfig: RpcConfig,
		Ticker:    ticker,
	}
}

func (s *BlockService) Start() error {
	clientL1 := jsonrpc.NewClient(s.RpcConfig.Url)

	for range s.Ticker.C {
		blockNumber, err := clientL1.EthBlockNumber()
		if err != nil {
			return fmt.Errorf("error getting most recent block: %w", err)
		}

		s.MostRecentL1Block = blockNumber
	}

	return nil
}

func (s *BlockService) Stop() error {
	s.Ticker.Stop()
	return nil
}

func (s *BlockService) GetMostRecentL1Block() uint64 {
	return s.MostRecentL1Block
}
