package rpc

import (
	"net/http"
)

// GetBlockchainStateResponse is the blockchain state RPC response
type GetBlockchainStateResponse struct {
	Success         bool            `json:"success"`
	BlockchainState BlockchainState `json:"blockchain_state"`
}

// BlockchainState blockchain state
type BlockchainState struct {
	Difficulty                  uint64      `json:"difficulty"`
	GenesisChallengeInitialized bool        `json:"genesis_challenge_initialized"`
	MempoolSize                 uint64      `json:"mempool_size"`
	Peak                        BlockRecord `json:"peak"`
//	Space                       uint64       `json:"space"` // @TODO this is supposed to be uint128 - will run out of room in uint64 pretty soon..
	SubSlotIters                uint64      `json:"sub_slot_iters"`
	Sync                        Sync        `json:"sync"`
}

// BlockRecord a single block record
type BlockRecord struct {
	ChallengeBlockInfoHash string            `json:"challenge_block_info_hash"`
	ChallengeVDFOutput     ClassgroupElement `json:"challenge_vdf_output"`
	Deficit                uint8             `json:"deficit"`
	FarmerPuzzleHash       string            `json:"farmer_puzzle_hash"`
	Fees                   uint64            `json:"fees"`
	//FinishedChallengeSlotHashes
	//FinishedInfusedChallengeSlotHashes
	//FinishedRewardSlotHashes
	HeaderHash                 string            `json:"header_hash"`
	Height                     uint32            `json:"height"`
	InfusedChallengeVDFOutput  ClassgroupElement `json:"infused_challenge_vdf_output"`
	Overflow                   bool              `json:"overflow"`
	PoolPuzzleHash             string            `json:"pool_puzzle_hash"`
	PrevHash                   string            `json:"prev_hash"`
	PrevTransactionBlockHash   string            `json:"prev_transaction_block_hash"`
	PrevTransactionBlockHeight uint32            `json:"prev_transaction_block_height"`
	RequiredIters              uint64            `json:"required_iters"`
	RewardClaimsIncorporated   []Coin            `json:"reward_claims_incorporated"`
	RewardInfusionNewChallenge string            `json:"reward_infusion_new_challenge"`
	SignagePointIndex          uint8             `json:"signage_point_index"`
	//SubEpochSummaryIncluded
	SubSlotIters uint64 `json:"sub_slot_iters"`
	Timestamp    uint64 `json:"timestamp"`
//	TotalIters   uint64 `json:"total_iters"` // @TODO this is supposed to be uint128
//	Weight       uint64 `json:"weight"`      // @TODO this is supposed to be uint128
}

// Coin is a coin // @TODO
type Coin struct{}

// ClassgroupElement Classgroup Element
type ClassgroupElement struct {
	Data string `json:"data"`
}

// Sync struct within blockchain state
type Sync struct {
	SyncMode           bool   `json:"sync_mode"`
	SyncProgressHeight uint32 `json:"sync_progress_height"`
	SyncTipHeight      uint32 `json:"sync_tip_height"`
	Synced             bool   `json:"synced"`
}

// GetBlockchainState returns blockchain state
func (c *Client) GetBlockchainState() (*GetBlockchainStateResponse, *http.Response, error) {
	request, err := c.NewRequest(http.MethodPost, ServiceFullNode, "get_blockchain_state", nil)
	if err != nil {
		return nil, nil, err
	}

	wbr := &GetBlockchainStateResponse{}
	resp, err := c.Do(request, wbr)
	if err != nil {
		return nil, resp, err
	}

	return wbr, resp, nil
}
