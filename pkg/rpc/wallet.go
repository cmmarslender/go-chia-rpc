package rpc

import "net/http"

// WalletService encapsulates wallet RPC methods
type WalletService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *WalletService) NewRequest(rpcEndpoint Endpoint, opt interface{}) (*Request, error) {
	return s.client.NewRequest(http.MethodPost, ServiceWallet, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *WalletService) Do(req *Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// CommonWalletOptions are common components to every wallet request
type CommonWalletOptions struct {
	WalletID uint32 `json:"wallet_id,omitempty"`
}

// GetWalletSyncStatusResponse Response for get_sync_status on wallet
type GetWalletSyncStatusResponse struct {
	GenesisInitialized bool `json:"genesis_initialized"`
	Success            bool `json:"success"`
	Synced             bool `json:"synced"`
	Syncing            bool `json:"syncing"`
}

// GetSyncStatus wallet rpc -> get_sync_status
func (s *WalletService) GetSyncStatus() (*GetWalletSyncStatusResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	wssr := &GetWalletSyncStatusResponse{}
	resp, err := s.Do(request, wssr)
	if err != nil {
		return nil, resp, err
	}

	return wssr, resp, nil
}

// GetWalletHeightInfoResponse response for get_height_info on wallet
type GetWalletHeightInfoResponse struct {
	Height  uint32 `json:"height"`
	Success bool   `json:"success"`
}

// GetHeightInfo wallet rpc -> get_height_info
func (s *WalletService) GetHeightInfo() (*GetWalletHeightInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletHeightInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletNetworkInfoResponse response for get_height_info on wallet
type GetWalletNetworkInfoResponse struct {
	NetworkName   string `json:"network_name"`
	NetworkPrefix string `json:"network_prefix"`
	Success       bool   `json:"success"`
}

// GetNetworkInfo wallet rpc -> get_network_info
func (s *WalletService) GetNetworkInfo() (*GetWalletNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletNetworkInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletsResponse wallet rpc -> get_wallets
type GetWalletsResponse struct {
	Success bool      `json:"success"`
	Wallets []*Wallet `json:"wallets"`
}

// WalletType types of wallets
type WalletType uint8

const (
	// WalletTypeStandard Standard Wallet
	WalletTypeStandard = WalletType(0)

	// WalletTypeRateLimited Rate Limited Wallet
	WalletTypeRateLimited = WalletType(1)

	// WalletTypeAtomicSwap Atomic Swap
	WalletTypeAtomicSwap = WalletType(2)

	// WalletTypeAuthorizedPayee Authorized Payee
	WalletTypeAuthorizedPayee = WalletType(3)

	// WalletTypeMultiSig Multi Sig
	WalletTypeMultiSig = WalletType(4)

	// WalletTypeCustody Custody
	WalletTypeCustody = WalletType(5)

	// WalletTypeColouredCoin Coloured Coin
	WalletTypeColouredCoin = WalletType(6)

	// WalletTypeRecoverable Recoverable Wallet
	WalletTypeRecoverable = WalletType(7)

	// WalletTypeDistributedID DID Wallet
	WalletTypeDistributedID = WalletType(8)
)

// Wallet single wallet record
type Wallet struct {
	Data string     `json:"data"`
	ID   uint32     `json:"id"`
	Name string     `json:"name"`
	Type WalletType `json:"type"`
}

// GetWallets wallet rpc -> get_wallets
func (s *WalletService) GetWallets() (*GetWalletsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletBalanceOptions request options for get_wallet_balance
type GetWalletBalanceOptions struct {
	CommonWalletOptions
}

// GetWalletBalanceResponse is the wallet balance RPC response
type GetWalletBalanceResponse struct {
	Success bool          `json:"success"`
	Balance WalletBalance `json:"wallet_balance"`
}

// WalletBalance specific wallet balance information
type WalletBalance struct {
	ConfirmedWalletBalance   int64 `json:"confirmed_wallet_balance"` // @TODO uint128
	MaxSendAmount            int64 `json:"max_send_amount"`
	PendingChange            int64 `json:"pending_change"`
	SpendableBalance         int64 `json:"spendable_balance"`          // @TODO uint128
	UnconfirmedWalletBalance int64 `json:"unconfirmed_wallet_balance"` // @TODO uint128
	UnspentCoinCount         int64 `json:"unspent_coin_count"`
	PendingCoinRemovalCount  int64 `json:"pending_coin_removal_count"`
	WalletID                 int32 `json:"wallet_id"`
}

// GetWalletBalance returns wallet balance
func (s *WalletService) GetWalletBalance(opts *GetWalletBalanceOptions) (*GetWalletBalanceResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	wbr := &GetWalletBalanceResponse{}
	resp, err := s.Do(request, wbr)
	if err != nil {
		return nil, resp, err
	}

	return wbr, resp, nil
}

// GetWalletTransactionsOptions options for get wallet transactions
type GetWalletTransactionsOptions struct {
	CommonWalletOptions
}

// GetWalletTransactionsResponse response for get_wallet_transactions
type GetWalletTransactionsResponse struct {
	Transactions []*Transaction `json:"transactions"`
	WalletID     uint32
}

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
	SpendBundle  SpendBundle     `json:"spend_bundle"`
	ToAddress    Address         `json:""`
	ToPuzzleHash PuzzleHash      `json:"to_puzzle_hash"`
	TradeID      string          `json:"trade_id"`
	Type         TransactionType `json:"type"`
	WalletID     uint32          `json:"wallet_id"`
}

// Address Own type for future methods to encode/decode
type Address string

// PuzzleHash Own type for future methods to encode/decode
type PuzzleHash string

// SentTo Represents the list of peers that we sent the transaction to, whether each one
// included it in the mempool, and what the error message (if any) was
// sent_to: List[Tuple[str, uint8, Optional[str]]]
// @TODO need to parse from the json
type SentTo struct {
	Peer                   string
	MempoolInclusionStatus MempoolInclusionStatus
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

// CoinSolution solution to a coin
type CoinSolution struct {
	Coin         Coin              `json:"coin"`
	PuzzleReveal SerializedProgram `json:"puzzle_reveal"`
	Solution     SerializedProgram `json:"solution"`
}

// SerializedProgram Just represent as a string for now
type SerializedProgram string

// GetTransactions wallet rpc -> get_transactions
func (s *WalletService) GetTransactions(opts *GetWalletTransactionsOptions) (*GetWalletTransactionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transactions", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
