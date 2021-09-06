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

// FullBlock a full block
type FullBlock struct {
	FinishedSubSlots             []*EndOfSubSlotBundle    `json:"finished_sub_slots"`
	RewardChainBlock             *RewardChainBlock        `json:"reward_chain_block"`
	ChallengeChainSPProof        *VDFProof                `json:"challenge_chain_sp_proof"`
	ChallengeChainIPProof        *VDFProof                `json:"challenge_chain_ip_proof"`
	RewardChainSPProof           *VDFProof                `json:"reward_chain_sp_proof"`
	RewardChainIPProof           *VDFProof                `json:"reward_chain_ip_proof"`
	InfusedChallengeChainIPProof *VDFProof                `json:"infused_challenge_chain_ip_proof"`
	Foliage                      *Foliage                 `json:"foliage"`
	FoliageTransactionBlock      *FoliageTransactionBlock `json:"foliage_transaction_block"`
	TransactionsInfo             *TransactionsInfo        `json:"transactions_info"`
	TransactionsGenerator        *SerializedProgram       `json:"transactions_generator"`          // @TODO Verify this is correct
	TransactionsGeneratorRefList []uint32                 `json:"transactions_generator_ref_list"` // @TODO Verify this is correct
}

// RewardChainBlock Reward Chain Block
type RewardChainBlock struct {
	Weight                     Uint128       `json:"weight"`
	Height                     uint32        `json:"height"`
	TotalIters                 Uint128       `json:"total_iters"`
	SignagePointIndex          uint8         `json:"signage_point_index"`
	POSSSCCChallengeHash       string        `json:"pos_ss_cc_challenge_hash"`
	ProofOfSpace               *ProofOfSpace `json:"proof_of_space"`
	ChallengeChainSPVDF        *VDFInfo      `json:"challenge_chain_sp_vdf"`
	ChallengeChainSPSignature  *G2Element    `json:"challenge_chain_sp_signature"`
	ChallengeChainIPVDF        *VDFInfo      `json:"challenge_chain_ip_vdf"`
	RewardChainSPVDF           *VDFInfo      `json:"reward_chain_sp_vdf"` // Not present for first sp in slot
	RewardChainSPSignature     *G2Element    `json:"reward_chain_sp_signature"`
	RewardChainIPVDF           *VDFInfo      `json:"reward_chain_ip_vdf"`
	InfusedChallengeChainIPVDF *VDFInfo      `json:"infused_challenge_chain_ip_vdf"` // Iff deficit < 16
	IsTransactionBlock         bool          `json:"is_transaction_block"`
}
