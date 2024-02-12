package chains

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/config"
	"time"
)

type Service struct {
	Config *config.ChainConfig
}

func NewService(Config *config.ChainConfig) *Service {
	return &Service{
		Config: Config,
	}
}

func (s Service) GetChains(context context.Context) ([]*Chain, error) {
	chains := createMockChains()
	return chains, nil
}

func createMockChains() []*Chain {
	return []*Chain{
		{
			Id:          1,
			NetworkName: "Ethereum",
			L1: &L1{
				ChainId:               "1",
				RpcUrl:                "https://mainnet.infura.io/v3/",
				RollupManagerAddress:  "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				RollupAddress:         "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				LatestL1BlockNumber:   12000000,
				HighestSequencedBlock: 11950000,
				HighestVerifiedBlock:  11900000,
			},
			L2: &L2{
				ChainId:           "100",
				DatastreamerUrl:   "https://data.ethereum.org",
				LatestBatchNumber: 5500,
				LatestBlockNumber: 11000,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
		{
			Id:          2,
			NetworkName: "Cardano",
			L1: &L1{
				ChainId:               "2",
				RpcUrl:                "https://rpc.cardano.org/",
				RollupManagerAddress:  "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				RollupAddress:         "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				LatestL1BlockNumber:   3000000,
				HighestSequencedBlock: 2998000,
				HighestVerifiedBlock:  2995000,
			},
			L2: &L2{
				ChainId:           "200",
				DatastreamerUrl:   "https://data.cardano.org",
				LatestBatchNumber: 2200,
				LatestBlockNumber: 4400,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
		{
			Id:          3,
			NetworkName: "Blueberry",
			L1: &L1{
				ChainId:               "3",
				RpcUrl:                "https://rpc.blueberry.net/",
				RollupManagerAddress:  "0xEfF10DB3c6445FB06411c6fc74fDCC8D1019aC7d",
				RollupAddress:         "0xEfF10DB3c6445FB06411c6fc74fDCC8D1019aC7d",
				LatestL1BlockNumber:   500000,
				HighestSequencedBlock: 499500,
				HighestVerifiedBlock:  499000,
			},
			L2: &L2{
				ChainId:           "300",
				DatastreamerUrl:   "https://data.blueberry.net",
				LatestBatchNumber: 1000,
				LatestBlockNumber: 2000,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
		{
			Id:          4,
			NetworkName: "Etrog",
			L1: &L1{
				ChainId:               "4",
				RpcUrl:                "https://rpc.etrog.com/",
				RollupManagerAddress:  "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				RollupAddress:         "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				LatestL1BlockNumber:   700000,
				HighestSequencedBlock: 699500,
				HighestVerifiedBlock:  699000,
			},
			L2: &L2{
				ChainId:           "400",
				DatastreamerUrl:   "https://data.etrog.com",
				LatestBatchNumber: 1500,
				LatestBlockNumber: 3000,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
		{
			Id:          5,
			NetworkName: "Another Chain",
			L1: &L1{
				ChainId:               "5",
				RpcUrl:                "https://rpc.etrog.com/",
				RollupManagerAddress:  "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				RollupAddress:         "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				LatestL1BlockNumber:   700000,
				HighestSequencedBlock: 699500,
				HighestVerifiedBlock:  699000,
			},
			L2: &L2{
				ChainId:           "400",
				DatastreamerUrl:   "https://data.etrog.com",
				LatestBatchNumber: 1500,
				LatestBlockNumber: 3000,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
		{
			Id:          6,
			NetworkName: "Another Chain",
			L1: &L1{
				ChainId:               "6",
				RpcUrl:                "https://rpc.etrog.com/",
				RollupManagerAddress:  "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				RollupAddress:         "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff",
				LatestL1BlockNumber:   700000,
				HighestSequencedBlock: 699500,
				HighestVerifiedBlock:  699000,
			},
			L2: &L2{
				ChainId:           "400",
				DatastreamerUrl:   "https://data.etrog.com",
				LatestBatchNumber: 1500,
				LatestBlockNumber: 3000,
				DatastreamStatus:  "Active",
			},
			LastUpdated: time.Now(),
		},
	}
}
