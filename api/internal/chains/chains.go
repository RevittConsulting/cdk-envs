package chains

import (
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
	"strconv"
	"time"
)

const (
	CardonaChain = "cardona"
)

type Chains struct {
	Config   *config.Chains
	Services *chain_services.Registry
}

func NewChains(Config *config.Chains, Services *chain_services.Registry) *Chains {
	return &Chains{
		Config:   Config,
		Services: Services,
	}
}

func (c Chains) getCardonaChain() (*Chain, error) {
	blockService := c.Services.GetService(chain_services.Logs)
	mostRecentL1Block := blockService.(*chain_services.LogsService).GetMostRecentL1Block()

	cardonaChain := &Chain{
		Id:          c.Config.Chains[CardonaChain].L1ChainId,
		NetworkName: "Polygon zkEVM Cardona Testnet",
		L1: &L1{
			ChainId:               strconv.Itoa(c.Config.Chains[CardonaChain].L1ChainId),
			RpcUrl:                c.Config.Chains[CardonaChain].L1RpcUrl,
			RollupManagerAddress:  c.Config.Chains[CardonaChain].RollupManagerAddress,
			RollupAddress:         c.Config.Chains[CardonaChain].RollupAddress,
			LatestL1BlockNumber:   int64(mostRecentL1Block),
			HighestSequencedBatch: 0, //
			HighestVerifiedBatch:  0, //
		},
		L2: &L2{
			ChainId:           strconv.Itoa(c.Config.Chains[CardonaChain].L2ChainId),
			DatastreamerUrl:   c.Config.Chains[CardonaChain].L2DataStreamUrl,
			LatestBatchNumber: 0,  //
			LatestBlockNumber: 0,  //
			DatastreamStatus:  "", //
		},
		LastUpdated: time.Now(),
	}

	return cardonaChain, nil
}
