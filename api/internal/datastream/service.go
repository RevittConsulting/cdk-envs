package datastream

import (
	"github.com/0xPolygonHermez/zkevm-data-streamer/datastreamer"
	"fmt"
)

type Config struct {
	Server     string
	StreamType datastreamer.StreamType
}

type Service struct {
	Config       *Config
	StreamClient *datastreamer.StreamClient
}

func NewService(cfg *Config) *Service {
	sc, err := datastreamer.NewClient(cfg.Server, cfg.StreamType)
	if err != nil {
		// TODO: import logger code to the project and use!
		fmt.Println("error creating datastreamer client")
		return nil
	}
	return &Service{
		Config:       cfg,
		StreamClient: sc,
	}
}

func (s *Service) GetTotalEntries() uint64 {
	total := s.StreamClient.GetTotalEntries()
	return total
}

/*
should stream from the last known l2 block using execCommandStartBookmark (s.StreamClient.ExecCommandStartBookmark(l2BlockNumber)) to get a streamClient from the 'l2blockno' we currently know of as highest from the l2, and then throw
away the stream as we get it (use erigon code to parse this stream, all we need is the latest l2 block number for now - display it and if it is equal to the rpc block number, we're good).
Given current setup this gives quite a complex dependency to satisfy. We could also pipe the stream down the ws in readable format...
*/

func (s *Service) GetByL2BlockNumber(l2BlockNumber uint64) {
	panic("implement me")
}
