package config

type Config struct {
	Port int

	Cardona CardonaConfig
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
