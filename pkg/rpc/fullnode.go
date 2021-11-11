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

	r := &GetBlockchainStateResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockOptions options for get_block rpc call
type GetBlockOptions struct {
	HeaderHash string `json:"header_hash"`
}

// GetBlockResponse response for get_block rpc call
type GetBlockResponse struct {
	Success bool             `json:"success"`
	Block   *types.FullBlock `json:"block"`
}

// GetBlock full_node->get_block RPC method
func (s *FullNodeService) GetBlock(opts *GetBlockOptions) (*GetBlockResponse, *http.Response, error) {
	request, err := s.NewRequest("get_block", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockByHeightOptions options for get_block_record_by_height and get_block rpc call
type GetBlockByHeightOptions struct {
	BlockHeight int `json:"height"`
}

type GetBlockRecordResponse struct {
	Success     bool               `json:"success"`
	BlockRecord *types.BlockRecord `json:"block_record"`
}

func (s *FullNodeService) GetBlockRecordByHeight(opts *GetBlockByHeightOptions) (*GetBlockRecordResponse, *http.Response, error) {
	// Get Block Record
	request, err := s.NewRequest("get_block_record_by_height", opts)
	if err != nil {
		return nil, nil, err
	}

	record := &GetBlockRecordResponse{}
	resp, err := s.Do(request, record)
	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
