package ws

type ChainData struct {
	MostRecentL1Block     uint64 `json:"mostRecentL1Block"`
	HighestSequencedBatch uint64 `json:"highestSequencedBatch"`
	MostRecentL2Batch     uint64 `json:"mostRecentL2Batch"`
	MostRecentL2Block     uint64 `json:"mostRecentL2Block"`
	DataStreamerStatus    string `json:"dataStreamerStatus"`
}
