package config

import (
	"github.com/RevittConsulting/cdk-envs/internal/buckets"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"github.com/RevittConsulting/cdk-envs/internal/tx"
)

type Config struct {
	Port         int
	DbFile       string
	ShutdownTime int

	RPC *jsonrpc.Config

	Chains  *Chains
	Buckets *buckets.Config
	Tx      *tx.Config

	L1Contracts *L1Contracts
}

type L1Contracts struct {
	SequencedBatchTopic         string
	VerificationTopic           string
	UpdateL1InfoTreeTopic       string
	InitialSequenceBatchesTopic string
}

type Chains struct {
	Chains map[string]*Chain
}

type Chain struct {
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
