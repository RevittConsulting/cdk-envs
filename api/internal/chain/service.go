package chain

import (
	"context"
	"encoding/json"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"io"
	"os"
)

type IChainDb interface {
	getHighestBlock(ctx context.Context) (*jsonrpc.Block, error)
}

type Service struct {
	Config *config.ChainConfig
	db     IChainDb
}

func NewService(db IChainDb, Config *config.ChainConfig) *Service {
	return &Service{
		Config: Config,
		db:     db,
	}
}

func (s Service) GetChains(context context.Context) (*Chain, error) {
	var results *Chain
	file, err := os.Open("mock_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s Service) GetHighestBlock(ctx context.Context) (*jsonrpc.Block, error) {
	return s.db.getHighestBlock(ctx)
}
