package ws

type ChainData struct {
	MostRecentL1Block uint64 `json:"mostRecentL1Block"`
	MostRecentL2Batch uint64 `json:"mostRecentL2Batch"`
}
