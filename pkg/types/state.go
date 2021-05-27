package types

// BlockchainState blockchain state
type BlockchainState struct {
	Difficulty                  uint64       `json:"difficulty"`
	GenesisChallengeInitialized bool         `json:"genesis_challenge_initialized"`
	MempoolSize                 uint64       `json:"mempool_size"`
	Peak                        *BlockRecord `json:"peak"`
	//Space                       uint64       `json:"space"` // @TODO this is supposed to be uint128 - will run out of room in uint64 pretty soon..
	SubSlotIters uint64 `json:"sub_slot_iters"`
	Sync         *Sync  `json:"sync"`
}

// Sync struct within blockchain state
type Sync struct {
	SyncMode           bool   `json:"sync_mode"`
	SyncProgressHeight uint32 `json:"sync_progress_height"`
	SyncTipHeight      uint32 `json:"sync_tip_height"`
	Synced             bool   `json:"synced"`
}
