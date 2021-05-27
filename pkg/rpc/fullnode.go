package rpc

import (
	"net/http"

	"github.com/cmmarslender/go-chia-rpc/pkg/types"
)

// FullNodeService encapsulates full node RPC methods
type FullNodeService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *FullNodeService) NewRequest(rpcEndpoint Endpoint, opt interface{}) (*Request, error) {
	return s.client.NewRequest(http.MethodPost, ServiceFullNode, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *FullNodeService) Do(req *Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetBlockchainStateResponse is the blockchain state RPC response
type GetBlockchainStateResponse struct {
	Success         bool                   `json:"success"`
	BlockchainState *types.BlockchainState `json:"blockchain_state"`
}

// GetBlockchainState returns blockchain state
func (s *FullNodeService) GetBlockchainState() (*GetBlockchainStateResponse, *http.Response, error) {
	request, err := s.NewRequest("get_blockchain_state", nil)
	if err != nil {
		return nil, nil, err
	}

	wbr := &GetBlockchainStateResponse{}
	resp, err := s.Do(request, wbr)
	if err != nil {
		return nil, resp, err
	}

	return wbr, resp, nil
}
