package rpc

import (
	"net/http"

	"github.com/cmmarslender/go-chia-rpc/pkg/types"
)

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
	Success bool            `json:"success"`
	Wallets []*types.Wallet `json:"wallets"`
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
	Success bool                 `json:"success"`
	Balance *types.WalletBalance `json:"wallet_balance"`
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
	Transactions []*types.Transaction `json:"transactions"`
	WalletID     uint32
}

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

// GetWalletTransactionOptions options for getting a single wallet transaction
type GetWalletTransactionOptions struct {
	CommonWalletOptions
	TransactionID string `json:"transaction_id"`
}

// GetWalletTransactionsResponse response for get_wallet_transactions
type GetWalletTransactionResponse struct {
	Transaction   *types.Transaction `json:"transaction"`
	TransactionID string             `json:"transaction_id"`
}

// GetTransaction returns a single transaction record
func (s *WalletService) GetTransaction(opts *GetWalletTransactionOptions) (*GetWalletTransactionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transaction", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
