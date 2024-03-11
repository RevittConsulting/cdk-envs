package chains

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
)

type HttpService struct {
	Config  *config.Chains
	Chains  *Chains
	Runtime *chain_services.Runtime
}

func NewService(config *config.Chains, services *chain_services.Registry, runtime *chain_services.Runtime) *HttpService {
	chains := NewChains(config, services)
	return &HttpService{
		Config:  config,
		Chains:  chains,
		Runtime: runtime,
	}
}

func (s *HttpService) GetChains(context context.Context) ([]*Chain, error) {
	chainNames, err := s.Chains.FindAllChainsFromConfig()
	if err != nil {
		return nil, err
	}

	chains, err := s.Chains.CreateChains(chainNames)
	if err != nil {
		return nil, err
	}

	return chains, nil
}

func (s *HttpService) ChangeChainService(context context.Context, chainName string) (string, error) {
	if err := s.Runtime.RestartService(chainName); err != nil {
		return "", err
	}

	return "service restarted", nil
}

// TODO: restart services creates an active service
// TODO: then fills out config for that service
// TODO: websocket to that card
// TODO: outline so we know its active

func (s *HttpService) StopServices(context context.Context) error {
	return s.Runtime.StopServices()
}
