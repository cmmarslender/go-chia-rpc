package types

// BlockRecord a single block record
type BlockRecord struct {
	ChallengeBlockInfoHash string             `json:"challenge_block_info_hash"`
	ChallengeVDFOutput     *ClassgroupElement `json:"challenge_vdf_output"`
	Deficit                uint8              `json:"deficit"`
	FarmerPuzzleHash       string             `json:"farmer_puzzle_hash"`
	Fees                   uint64             `json:"fees"`
	//FinishedChallengeSlotHashes
	//FinishedInfusedChallengeSlotHashes
	//FinishedRewardSlotHashes
	HeaderHash                 string             `json:"header_hash"`
	Height                     uint32             `json:"height"`
	InfusedChallengeVDFOutput  *ClassgroupElement `json:"infused_challenge_vdf_output"`
	Overflow                   bool               `json:"overflow"`
	PoolPuzzleHash             string             `json:"pool_puzzle_hash"`
	PrevHash                   string             `json:"prev_hash"`
	PrevTransactionBlockHash   string             `json:"prev_transaction_block_hash"`
	PrevTransactionBlockHeight uint32             `json:"prev_transaction_block_height"`
	RequiredIters              uint64             `json:"required_iters"`
	RewardClaimsIncorporated   []*Coin            `json:"reward_claims_incorporated"`
	RewardInfusionNewChallenge string             `json:"reward_infusion_new_challenge"`
	SignagePointIndex          uint8              `json:"signage_point_index"`
	//SubEpochSummaryIncluded
	SubSlotIters uint64  `json:"sub_slot_iters"`
	Timestamp    uint64  `json:"timestamp"`
	TotalIters   Uint128 `json:"total_iters"`
	Weight       Uint128 `json:"weight"`
}
