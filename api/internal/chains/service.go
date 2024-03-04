package chains

import (
	"context"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/chains/chain_services"
)

type Service struct {
	Config *config.Chains
	Chains *Chains
}

func NewService(Config *config.Chains, Services *chain_services.Registry) *Service {
	chains := NewChains(Config, Services)
	return &Service{
		Config: Config,
		Chains: chains,
	}
}

func (s Service) GetChains(context context.Context) ([]*Chain, error) {
	chains := make([]*Chain, 0)

	cardona, err := s.Chains.getCardonaChain()
	if err != nil {
		return nil, err
	}
	chains = append(chains, cardona)

	return chains, nil
}
