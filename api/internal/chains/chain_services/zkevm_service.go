package chain_services

import (
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"log"
	"net"
	"time"
)

var (
	DataStreamerStatusDown = "down"
	DataStreamerStatusUp   = "up"
)

type ZkEvmService struct {
	Config    *config.Chains
	RpcConfig *config.RPCConfig
	Ticker    *time.Ticker

	stopChan chan struct{}

	MostRecentL2Block  uint64
	MostRecentL2Batch  uint64
	DataStreamerStatus string
}

func NewZkEvmService(Config *config.Chains, RpcConfig *config.RPCConfig) *ZkEvmService {
	ticker := time.NewTicker(5 * time.Second)
	return &ZkEvmService{
		Config:    Config,
		RpcConfig: RpcConfig,
		Ticker:    ticker,
		stopChan:  make(chan struct{}),
	}
}

func (s *ZkEvmService) Start() error {
	s.Ticker = time.NewTicker(5 * time.Second)
	s.stopChan = make(chan struct{})

	clientL2 := jsonrpc.NewClient(s.Config.Chains[ActiveChainConfigName].L2RpcUrl)

	log.Println("zkevm service started")
	go func() {
		for {
			select {
			case <-s.Ticker.C:
				batchNum, err := clientL2.ZkGetBatchNumber()
				if err != nil {
					log.Println("error getting most recent batch")
				}
				s.MostRecentL2Batch = batchNum

				blockNum, err := clientL2.EthBlockNumber()
				if err != nil {
					log.Println("error getting most recent block")
				}
				s.MostRecentL2Block = blockNum

				dsStatus, err := checkDataStreamerStatus(s.Config.Chains[ActiveChainConfigName].L2DataStreamUrl)
				if err != nil {
					log.Println("error getting data streamer status")
				}
				s.DataStreamerStatus = dsStatus
			case <-s.stopChan:
				return
			}
		}
	}()

	return nil
}

func (s *ZkEvmService) Stop() error {
	log.Println("zkevm service stopped")

	if s.stopChan != nil {
		close(s.stopChan)
		s.stopChan = nil
	}

	if s.Ticker != nil {
		s.Ticker.Stop()
		s.Ticker = nil
	}

	return nil
}

func (s *ZkEvmService) GetMostRecentL2Batch() uint64 {
	return s.MostRecentL2Batch
}

func (s *ZkEvmService) GetMostRecentL2Block() uint64 {
	return s.MostRecentL2Block
}

func (s *ZkEvmService) GetDataStreamerStatus() string {
	return s.DataStreamerStatus
}

func checkDataStreamerStatus(url string) (string, error) {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		return DataStreamerStatusDown, err
	}
	defer conn.Close()
	return DataStreamerStatusUp, nil
}
