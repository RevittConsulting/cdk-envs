package tx

import (
	"context"
	"math/big"
)

type HttpService struct {
	Config *Config
}

func NewService(config *Config) *HttpService {
	return &HttpService{
		Config: config,
	}
}

func (s *HttpService) CreateTx(ctx context.Context, req *Request) (*Response, error) {
	res := &Response{}
	input := &Input{
		Host:        req.Host,
		ToAddress:   req.ToAddress,
		FromAddress: req.FromAddress,
		PrivateKey:  req.PrivateKey,
		Amount:      big.NewInt(int64(req.Amount)),
		GasLimit:    uint64(req.GasLimit),
		GasPrice:    big.NewInt(int64(req.GasPrice)),
	}

	tx, err := s.DoTx(ctx, input)
	if err != nil {
		res.Error = err.Error()
		return res, nil
	}

	res.Output = tx
	return res, nil
}
