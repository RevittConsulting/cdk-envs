package config

import "time"

type Config struct {
	Port int

	RPC     RPCConfig
	Cardona CardonaConfig
	Chain   ChainConfig
}

type RPCConfig struct {
	Url  string
	Url2 string

	PollingInterval time.Duration
}

type CardonaConfig struct {
	L1ChainId int
	L1RpcUrl  string

	L2ChainId       int
	L2RpcUrl        string
	L2DataStreamUrl string

	RollupManagerAddress string
	RollupAddress        string

	TopicsVerification string
	TopicsSequence     string
}

type ChainConfig struct {
}
