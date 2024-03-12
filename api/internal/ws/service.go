package ws

import (
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
)

type Service struct {
	Services *chain_services.Runtime
}

func NewService(Services *chain_services.Runtime) *Service {
	return &Service{
		Services: Services,
	}
}

func (s *Service) PollChainData() (*ChainData, error) {
	chainData := &ChainData{}

	activeServices := s.Services.GetActiveServices()
	if activeServices == nil {
		return chainData, nil
	}

	for _, service := range activeServices {
		switch v := service.(type) {
		case *chain_services.LogsService:
			chainData.MostRecentL1Block = v.GetMostRecentL1Block()
			chainData.HighestSequencedBatch = v.GetHighestSequencedBatch()
			chainData.HighestVerifiedBatch = v.GetHighestVerifiedBatch()
		case *chain_services.ZkEvmService:
			chainData.MostRecentL2Batch = v.GetMostRecentL2Batch()
			chainData.MostRecentL2Block = v.GetMostRecentL2Block()
			chainData.DataStreamerStatus = v.GetDataStreamerStatus()
		}
	}

	return chainData, nil
}
