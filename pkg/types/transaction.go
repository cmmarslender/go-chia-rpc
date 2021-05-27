package types

// Transaction Single Transaction
type Transaction struct {
	Additions         []*Coin `json:"additions"`
	Amount            uint64  `json:"amount"`
	Confirmed         bool    `json:"confirmed"`
	ConfirmedAtHeight uint32  `json:"confirmed_at_height"`
	CreatedAtTime     uint64  `json:"created_at_time"`
	FeeAmount         uint64  `json:"fee_amount"`
	Name              string  `json:"name"` // @TODO bytes32 / hex
	Removals          []*Coin `json:"removals"`
	Sent              uint32  `json:"sent"`
	//SentTo            SentTo          `json:"sent_to"` // @TODO need to properly unserialize this
	SpendBundle  *SpendBundle     `json:"spend_bundle"`
	ToAddress    *Address         `json:""`
	ToPuzzleHash *PuzzleHash      `json:"to_puzzle_hash"`
	TradeID      string           `json:"trade_id"`
	Type         *TransactionType `json:"type"`
	WalletID     uint32           `json:"wallet_id"`
}

// Address Own type for future methods to encode/decode
type Address string

// SentTo Represents the list of peers that we sent the transaction to, whether each one
// included it in the mempool, and what the error message (if any) was
// sent_to: List[Tuple[str, uint8, Optional[str]]]
// @TODO need to parse from the json
type SentTo struct {
	Peer                   string
	MempoolInclusionStatus *MempoolInclusionStatus
	Error                  string
}

// MempoolInclusionStatus status of being included in the mempool
type MempoolInclusionStatus uint8

const (
	// MempoolInclusionStatusSuccess Successfully added to mempool
	MempoolInclusionStatusSuccess = MempoolInclusionStatus(1)

	// MempoolInclusionStatusPending Pending being added to the mempool
	MempoolInclusionStatusPending = MempoolInclusionStatus(2)

	// MempoolInclusionStatusFailed Failed being added to the mempool
	MempoolInclusionStatusFailed = MempoolInclusionStatus(3)
)

// TransactionType type of transaction
type TransactionType uint32

const (
	// TransactionTypeIncomingTX incoming transaction
	TransactionTypeIncomingTX = TransactionType(0)

	// TransactionTypeOutgoingTX outgoing transaction
	TransactionTypeOutgoingTX = TransactionType(1)

	// TransactionTypeCoinbaseReward coinbase reward
	TransactionTypeCoinbaseReward = TransactionType(2)

	// TransactionTypeFeeReward fee reward
	TransactionTypeFeeReward = TransactionType(3)

	// TransactionTypeIncomingTrade incoming trade
	TransactionTypeIncomingTrade = TransactionType(4)

	// TransactionTypeOutgoingTrade outgoing trade
	TransactionTypeOutgoingTrade = TransactionType(5)
)

// SpendBundle Spend Bundle...
type SpendBundle struct {
	AggregatedSignature string          `json:"aggregated_signature"`
	CoinSolutions       []*CoinSolution `json:"coin_solutions"`
}
