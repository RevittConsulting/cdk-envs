package chains

import (
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
	"strconv"
	"time"
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

func (c *Chains) FindAllChainsFromConfig() ([]string, error) {
	chains := make([]string, 0)
	for k := range c.Config.Chains {
		chains = append(chains, k)
	}

	return chains, nil
}

func (c *Chains) CreateChains(chainNames []string) ([]*Chain, error) {
	//blockService := c.Services.GetService(chain_services.Logs)
	//mostRecentL1Block := blockService.(*chain_services.LogsService).GetMostRecentL1Block()

	chains := make([]*Chain, 0)
	for _, serviceName := range chainNames {
		chainConfig := c.Config.Chains[serviceName]
		chain := &Chain{
			ServiceName: serviceName,
			NetworkName: chainConfig.NetworkName,
			L1: &L1{
				ChainId:               strconv.Itoa(chainConfig.L1ChainId),
				RpcUrl:                chainConfig.L1RpcUrl,
				RollupManagerAddress:  chainConfig.RollupManagerAddress,
				RollupAddress:         chainConfig.RollupAddress,
				LatestL1BlockNumber:   0,
				HighestSequencedBatch: 0,
				HighestVerifiedBatch:  0,
			},
			L2: &L2{
				ChainId:           strconv.Itoa(chainConfig.L2ChainId),
				DatastreamerUrl:   chainConfig.L2DataStreamUrl,
				LatestBatchNumber: 0,
				LatestBlockNumber: 0,
				DatastreamStatus:  "unknown",
			},
			LastUpdated: time.Now(),
		}
		chains = append(chains, chain)
	}

	return chains, nil
}
