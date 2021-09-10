package rpc

import (
	"crypto/tls"
	"time"
)

// ClientOptionFunc can be used to customize a new RPC client.
type ClientOptionFunc func(*Client) error

// WithBaseURL sets the host for RPC requests
func WithBaseURL(url string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(url)
	}
}

// WithNodePort sets the port for full node RPC communication
func WithNodePort(port uint16) ClientOptionFunc {
	return func(c *Client) error {
		c.nodePort = port
		return nil
	}
}

// WithNodeKeyPair specify keypair to use with full node RPC requests
func WithNodeKeyPair(keyPair *tls.Certificate) ClientOptionFunc {
	return func(c *Client) error {
		c.nodeKeyPair = keyPair
		return nil
	}
}

// WithFarmerPort sets the port for farmer RPC communication
func WithFarmerPort(port uint16) ClientOptionFunc {
	return func(c *Client) error {
		c.farmerPort = port
		return nil
	}
}

// WithFarmerKeyPair specify keypair to use with farmer RPC requests
func WithFarmerKeyPair(keyPair *tls.Certificate) ClientOptionFunc {
	return func(c *Client) error {
		c.farmerKeyPair = keyPair
		return nil
	}
}

// WithHarvesterPort sets the port for harvester RPC communication
func WithHarvesterPort(port uint16) ClientOptionFunc {
	return func(c *Client) error {
		c.harvesterPort = port
		return nil
	}
}

// WithHarvesterKeyPair specify keypair to use with harvester RPC requests
func WithHarvesterKeyPair(keyPair *tls.Certificate) ClientOptionFunc {
	return func(c *Client) error {
		c.harvesterKeyPair = keyPair
		return nil
	}
}

// WithWalletPort sets the port for wallet RPC communication
func WithWalletPort(port uint16) ClientOptionFunc {
	return func(c *Client) error {
		c.walletPort = port
		return nil
	}
}

// WithWalletKeyPair specify keypair to use with wallet RPC requests
func WithWalletKeyPair(keyPair *tls.Certificate) ClientOptionFunc {
	return func(c *Client) error {
		c.walletKeyPair = keyPair
		return nil
	}
}

// WithCache specify a duration http requests should be cached for
// If unset, cache will not be used
func WithCache(validTime time.Duration) ClientOptionFunc {
	return func(c *Client) error {
		c.CacheValidTime = validTime

		return nil
	}
}
