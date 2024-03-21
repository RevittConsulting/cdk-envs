package chain_services

import (
	"fmt"
	"github.com/RevittConsulting/chain-dev-utils/internal/jsonrpc"
	"time"
)

type BlockService struct {
	RpcConfig *jsonrpc.Config
	Ticker    *time.Ticker

	MostRecentL1Block uint64
}

func NewBlockService(RpcConfig *jsonrpc.Config) *BlockService {
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
