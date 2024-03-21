package chains

import "time"

type Chain struct {
	NetworkName string `json:"networkName" db:"network_name"`
	L1          *L1
	L2          *L2
	LastUpdated time.Time `json:"lastUpdated" db:"last_updated"`
	ServiceName string    `json:"serviceName" db:"service_name"`
}

type L1 struct {
	ChainId               string `json:"chainId" db:"chain_id"`
	RpcUrl                string `json:"rpcUrl" db:"rpc_url"`
	Etherscan             string `json:"etherscan" db:"etherscan"`
	Blockscout            string `json:"blockscout" db:"blockscout"`
	RollupManagerAddress  string `json:"rollupManagerAddress" db:"rollup_manager_address"`
	RollupAddress         string `json:"rollupAddress" db:"rollup_address"`
	LatestL1BlockNumber   int64  `json:"latestL1BlockNumber" db:"latest_l1_block_number"`
	HighestSequencedBatch int64  `json:"highestSequencedBatch" db:"highest_sequenced_batch"`
	HighestVerifiedBatch  int64  `json:"highestVerifiedBatch" db:"highest_verified_batch"`
}

type L2 struct {
	ChainId           string `json:"chainId" db:"chain_id"`
	RpcUrl            string `json:"rpcUrl" db:"rpc_url"`
	Polygonscan       string `json:"polygonscan" db:"polygonscan"`
	DatastreamerUrl   string `json:"datastreamerUrl" db:"datastreamer_url"`
	LatestBatchNumber int64  `json:"latestBatchNumber" db:"latest_batch_number"`
	LatestBlockNumber int64  `json:"latestBlockNumber" db:"latest_block_number"`
	DatastreamStatus  string `json:"datastreamStatus" db:"datastream_status"`
}

type ChainRequest struct {
	ServiceName string `json:"serviceName"`
}
