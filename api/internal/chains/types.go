package chains

import "time"

type Chain struct {
	Id          int    `json:"id" db:"id"`
	NetworkName string `json:"networkName" db:"network_name"`
	L1          *L1
	L2          *L2
	LastUpdated time.Time `json:"lastUpdated" db:"last_updated"`
}

type L1 struct {
	ChainId               string `json:"chainId" db:"chain_id"`
	RpcUrl                string `json:"rpcUrl" db:"rpc_url"`
	RollupManagerAddress  string `json:"rollupManagerAddress" db:"rollup_manager_address"`
	RollupAddress         string `json:"rollupAddress" db:"rollup_address"`
	LatestL1BlockNumber   int64  `json:"latestL1BlockNumber" db:"latest_l1_block_number"`
	HighestSequencedBlock int64  `json:"highestSequencedBlock" db:"highest_sequenced_block"`
	HighestVerifiedBlock  int64  `json:"highestVerifiedBlock" db:"highest_verified_block"`
}

type L2 struct {
	ChainId           string `json:"chainId" db:"chain_id"`
	DatastreamerUrl   string `json:"datastreamerUrl" db:"datastreamer_url"`
	LatestBatchNumber int64  `json:"latestBatchNumber" db:"latest_batch_number"`
	LatestBlockNumber int64  `json:"latestBlockNumber" db:"latest_block_number"`
	DatastreamStatus  string `json:"datastreamStatus" db:"datastream_status"`
}
