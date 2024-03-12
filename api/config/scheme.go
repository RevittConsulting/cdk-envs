package config

import "time"

type Config struct {
	Port         int
	DbFile       string
	ShutdownTime int

	RPC *RPCConfig

	Chains  *Chains
	Buckets *BucketsConfig

	L1Contracts *L1Contracts
}

type L1Contracts struct {
	SequencedBatchTopic         string
	VerificationTopic           string
	UpdateL1InfoTreeTopic       string
	InitialSequenceBatchesTopic string
}

type RPCConfig struct {
	Url        string
	Url2       string
	ZkEvm      string
	CardonaUrl string

	PollingInterval time.Duration
}

type Chains struct {
	Chains map[string]*ChainConfig
}

type ChainConfig struct {
	NetworkName string

	Etherscan      string
	Blockscout     string
	Polygonscan    string
	CurrencySymbol string

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

type BucketsConfig struct{}
