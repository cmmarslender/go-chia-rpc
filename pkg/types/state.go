package types

// WebsocketBlockchainState is how blockchain_state comes across in the websockets (wrapped)
type WebsocketBlockchainState struct {
	BlockchainState BlockchainState `json:"blockchain_state"`
}

// BlockchainState blockchain state
type BlockchainState struct {
	Difficulty                  uint64       `json:"difficulty"`
	GenesisChallengeInitialized bool         `json:"genesis_challenge_initialized"`
	MempoolSize                 uint64       `json:"mempool_size"`
	MempoolCost                 uint64       `json:"mempool_cost"`
	Peak                        *BlockRecord `json:"peak"`
	Space                       Uint128      `json:"space"`
	SubSlotIters                uint64       `json:"sub_slot_iters"`
	Sync                        *Sync        `json:"sync"`
	CompactBlocks               uint32       `json:"compact_blocks"`
	UncompactBlocks             uint32       `json:"uncompact_blocks"`
	HintCount                   uint64       `json:"hint_count"`
}

// Sync struct within blockchain state
type Sync struct {
	SyncMode           bool   `json:"sync_mode"`
	SyncProgressHeight uint32 `json:"sync_progress_height"`
	SyncTipHeight      uint32 `json:"sync_tip_height"`
	Synced             bool   `json:"synced"`
}
