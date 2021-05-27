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
	WalletID uint64 `json:"wallet_id,omitempty"`
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
