package rpc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cmmarslender/go-chia-rpc/pkg/config"
	"github.com/google/go-querystring/query"
	"github.com/gorilla/websocket"
)

const (
	defaultHost string = "localhost"

	// ServiceDaemon the websocket/daemon service
	ServiceDaemon ServiceType = 0

	// ServiceFullNode the full node service
	ServiceFullNode ServiceType = 1

	// ServiceFarmer the farmer service
	ServiceFarmer ServiceType = 2

	// ServiceHarvester the harvester service
	ServiceHarvester ServiceType = 3

	// ServiceWallet the wallet service
	ServiceWallet ServiceType = 4
)

// ServiceType is a type that refers to a particular service
type ServiceType uint8

// Endpoint represents and RPC Method
type Endpoint string

// Client is the RPC client
type Client struct {
	config  *config.ChiaConfig
	baseURL *url.URL

	// If set > 0, will configure http requests with a cache
	CacheValidTime time.Duration

	daemonPort    uint16
	daemonKeyPair *tls.Certificate
	daemonDialer  *websocket.Dialer

	nodePort    uint16
	nodeKeyPair *tls.Certificate
	nodeClient  *http.Client

	farmerPort    uint16
	farmerKeyPair *tls.Certificate
	farmerClient  *http.Client

	harvesterPort    uint16
	harvesterKeyPair *tls.Certificate
	harvesterClient  *http.Client

	walletPort    uint16
	walletKeyPair *tls.Certificate
	walletClient  *http.Client

	// Services for the different chia services
	DaemonService   *DaemonService
	FullNodeService *FullNodeService
	WalletService   *WalletService
}

// NewClient returns a new RPC Client
func NewClient(options ...ClientOptionFunc) (*Client, error) {
	cfg, err := config.GetChiaConfig()
	if err != nil {
		return nil, err
	}

	c := &Client{
		config:        cfg,
		daemonPort:    cfg.DaemonPort,
		nodePort:      cfg.FullNode.RPCPort,
		farmerPort:    cfg.Farmer.RPCPort,
		harvesterPort: cfg.Harvester.RPCPort,
		walletPort:    cfg.Wallet.RPCPort,
	}

	err = c.setBaseURL(&url.URL{
		Scheme: "https",
		Host:   defaultHost,
	})
	if err != nil {
		return nil, err
	}

	err = c.initialKeyPairs()
	if err != nil {
		return nil, err
	}

	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	// Generate the http clients and transports after any client options are applied, in case custom keypairs were provided
	err = c.generateHTTPClients()
	if err != nil {
		return nil, err
	}

	// Init Services
	c.DaemonService = &DaemonService{client: c}
	c.FullNodeService = &FullNodeService{client: c}
	c.WalletService = &WalletService{client: c}

	return c, nil
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(url *url.URL) error {
	c.baseURL = url

	return nil
}

// Sets the initial key pairs based on config
func (c *Client) initialKeyPairs() error {
	var err error

	c.daemonKeyPair, err = c.config.DaemonSSL.LoadPrivateKeyPair()
	if err != nil {
		return err
	}

	c.nodeKeyPair, err = c.config.FullNode.SSL.LoadPrivateKeyPair()
	if err != nil {
		return err
	}

	c.farmerKeyPair, err = c.config.Farmer.SSL.LoadPrivateKeyPair()
	if err != nil {
		return err
	}

	c.harvesterKeyPair, err = c.config.Harvester.SSL.LoadPrivateKeyPair()
	if err != nil {
		return err
	}

	c.walletKeyPair, err = c.config.Wallet.SSL.LoadPrivateKeyPair()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) generateHTTPClients() error {
	var err error

	if c.daemonDialer == nil {
		c.daemonDialer = &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{*c.daemonKeyPair},
				InsecureSkipVerify: true,
			},
		}
	}

	if c.nodeClient == nil {
		c.nodeClient, err = c.generateHTTPClientForService(ServiceFullNode)
		if err != nil {
			return err
		}
	}

	if c.farmerClient == nil {
		c.farmerClient, err = c.generateHTTPClientForService(ServiceFarmer)
		if err != nil {
			return err
		}
	}

	if c.harvesterClient == nil {
		c.harvesterClient, err = c.generateHTTPClientForService(ServiceHarvester)
		if err != nil {
			return err
		}
	}

	if c.walletClient == nil {
		c.walletClient, err = c.generateHTTPClientForService(ServiceWallet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) generateHTTPClientForService(service ServiceType) (*http.Client, error) {
	var keyPair *tls.Certificate

	switch service {
	case ServiceFullNode:
		keyPair = c.nodeKeyPair
	case ServiceFarmer:
		keyPair = c.farmerKeyPair
	case ServiceHarvester:
		keyPair = c.harvesterKeyPair
	case ServiceWallet:
		keyPair = c.walletKeyPair
	default:
		return nil, fmt.Errorf("unknown service")
	}

	var transport http.RoundTripper

	transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{*keyPair},
			InsecureSkipVerify: true, // Cert is apparently for chia.net - can't validate until it matches hostname
		},
	}

	if c.CacheValidTime > 0 {
		transport = NewCachedTransport(c.CacheValidTime, transport)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	return client, nil
}

func (c *Client) portForService(service ServiceType) uint16 {
	var port uint16 = 0

	switch service {
	case ServiceFullNode:
		port = c.nodePort
	case ServiceFarmer:
		port = c.farmerPort
	case ServiceHarvester:
		port = c.harvesterPort
	case ServiceWallet:
		port = c.walletPort
	}

	return port
}

// HTTPClientForService returns the proper http client to use with the service
func (c *Client) HTTPClientForService(service ServiceType) (*http.Client, error) {
	var client *http.Client

	switch service {
	case ServiceFullNode:
		client = c.nodeClient
	case ServiceFarmer:
		client = c.farmerClient
	case ServiceHarvester:
		client = c.harvesterClient
	case ServiceWallet:
		client = c.walletClient
	}

	if client == nil {
		return nil, fmt.Errorf("unknown service")
	}

	return client, nil
}

// Request is a wrapped http.Request that indicates the service we're making the RPC call to
type Request struct {
	Service ServiceType
	Request *http.Request
}

// NewRequest creates an RPC request for the specified service
func (c *Client) NewRequest(method string, service ServiceType, rpcEndpoint Endpoint, opt interface{}) (*Request, error) {
	u := *c.baseURL

	u.Host = fmt.Sprintf("%s:%d", u.Host, c.portForService(service))

	u.RawPath = fmt.Sprintf("/%s", rpcEndpoint)
	u.Path = fmt.Sprintf("/%s", rpcEndpoint)

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	var err error
	switch {
	case method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		// Always need at least an empty json object in the body
		if opt == nil {
			opt = []byte(`{}`)
		}

		body, err = json.Marshal(opt)
		if err != nil {
			return nil, err
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return &Request{
		Service: service,
		Request: req,
	}, nil
}

// Do sends an RPC request and returns the RPC response.
func (c *Client) Do(req *Request, v interface{}) (*http.Response, error) {
	client, err := c.HTTPClientForService(req.Service)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}
