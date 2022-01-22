package types

// UnfinishedBlock data from unfinished_block websocket event
type UnfinishedBlock struct {
	MaxCost   uint64 `json:"max_cost"`
	BlockCost uint64 `json:"block_cost"`
	BlockFees uint64 `json:"block_fees"`
}
