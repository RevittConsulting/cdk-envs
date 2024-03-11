package ws

import (
	"encoding/json"
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

func (s *Service) PollChainData() ([]byte, error) {
	activeServices := s.Services.GetActiveServices()
	if activeServices == nil {
		return []byte{}, nil
	}

	chainData := &ChainData{}
	for _, service := range activeServices {
		switch v := service.(type) {
		case *chain_services.LogsService:
			chainData.MostRecentL1Block = v.GetMostRecentL1Block()
		}
	}

	bytes, err := json.Marshal(chainData)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
