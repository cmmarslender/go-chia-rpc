package rpc

import "net/http"

// WalletService encapsulates wallet RPC methods
type WalletService struct {
	client *Client
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
	request, err := s.client.NewRequest(http.MethodPost, ServiceWallet, "get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	wssr := &GetWalletSyncStatusResponse{}
	resp, err := s.client.Do(request, wssr)
	if err != nil {
		return nil, resp, err
	}

	return wssr, resp, nil
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
	request, err := s.client.NewRequest(http.MethodPost, ServiceWallet, "get_wallet_balance", opts)
	if err != nil {
		return nil, nil, err
	}

	wbr := &GetWalletBalanceResponse{}
	resp, err := s.client.Do(request, wbr)
	if err != nil {
		return nil, resp, err
	}

	return wbr, resp, nil
}
