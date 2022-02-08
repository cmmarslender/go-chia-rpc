package rpc

import (
	"net/http"

	"github.com/cmmarslender/go-chia-rpc/pkg/rpcinterface"
	"github.com/cmmarslender/go-chia-rpc/pkg/types"
)

// CrawlerService encapsulates crawler RPC methods
type CrawlerService struct {
	client *Client
}

// NewRequest returns a new request specific to the crawler service
func (s *CrawlerService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceCrawler, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *CrawlerService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetPeerCountsResponse Response for get_get_peer_counts on crawler
type GetPeerCountsResponse struct {
	Success    bool                     `json:"success"`
	PeerCounts *types.CrawlerPeerCounts `json:"peer_counts"`
}

// GetPeerCounts crawler rpc -> get_peer_counts
func (s *CrawlerService) GetPeerCounts() (*GetPeerCountsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_peer_counts", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetPeerCountsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
