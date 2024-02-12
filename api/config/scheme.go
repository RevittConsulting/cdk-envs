package config

import "time"

type Config struct {
	Port   int
	DbFile string

	RPC     RPCConfig
	Cardona CardonaConfig
	Chain   ChainConfig
	Buckets BucketsConfig
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

type ChainConfig struct{}

type BucketsConfig struct{}
