package ws

import "github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"

type Service struct {
	Services *chain_services.Runtime
}

func NewService(Services *chain_services.Runtime) *Service {
	return &Service{
		Services: Services,
	}
}

func (s *Service) PollChainData() (uint64, error) {
	activeService := s.Services.GetActiveService()
	if activeService == nil {
		return 0, nil
	}

	mostRecentL1Block := activeService.(*chain_services.LogsService).GetMostRecentL1Block()
	return mostRecentL1Block, nil
}
