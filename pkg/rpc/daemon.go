package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/cmmarslender/go-chia-rpc/pkg/types"
	"github.com/gorilla/websocket"
	"net/url"
)

const origin string = "go-chia-rpc"

type websocketRespHandler func(*types.WebsocketResponse, error)

// DaemonService encapsulates websocket functionality with the daemon
type DaemonService struct {
	client   *Client
	conn     *websocket.Conn
	handlers []websocketRespHandler
}

// ensureConnection ensures there is an open websocket connection
func (d *DaemonService) ensureConnection() error {
	if d.conn == nil {
		u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s:%d", d.client.baseURL.Host, d.client.daemonPort), Path: "/"}
		var err error
		d.conn, _, err = d.client.daemonDialer.Dial(u.String(), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// CloseConnection closes the websocket connection if it is open
func (d *DaemonService) CloseConnection() error {
	if d.conn != nil {
		d.conn = nil

		err := d.conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Do sends a request over the websocket connection
func (d *DaemonService) Do(req *types.WebsocketRequest) error {
	err := d.ensureConnection()
	if err != nil {
		return err
	}

	return d.conn.WriteJSON(req)
}

// AddHandler adds an additional handler function to call when a message is received over the websocket
// This is expected to NOT be used in conjunction with ListenSync
// This will run in the background, and allow other things to happen in the foreground
// while ListenSync will take over the foreground process
func (d *DaemonService) AddHandler(handler websocketRespHandler) error {
	d.handlers = append(d.handlers, handler)

	go d.ListenSync(d.handlerProxy)
	return nil
}

// handlerProxy matches the websocketRespHandler signature to send requests back to any registered handlers
// Here to support multiple handlers for a single event in the future
func (d *DaemonService) handlerProxy(resp *types.WebsocketResponse, err error) {
	for _, handler := range d.handlers {
		handler(resp, err)
	}
}

// ListenSync Listens for responses over the websocket connection in the foreground
// The error returned from this function would only correspond to an error setting up the listener
// Errors returned by ReadMessage, or some other part of the websocket request/response will be
// passed to the handler to deal with
func (d *DaemonService) ListenSync(handler websocketRespHandler) error {
	for {
		_, message, err := d.conn.ReadMessage()
		resp := &types.WebsocketResponse{}
		err = json.Unmarshal(message, resp)
		handler(resp, err)
	}
}

// SubscribeSelf calls subscribe for any requests that this client makes to the server
// Different from Subscribe with a custom service - that is more for subscribing to built in events emitted by Chia
// This call will subscribe `go-chia-rpc` origin for any requests we specifically make of the server
func (d *DaemonService) SubscribeSelf() error {
	return d.Subscribe(origin)
}

// Subscribe adds a subscription to a particular service
func (d *DaemonService) Subscribe(service string) error {
	request := &types.WebsocketRequest{
		Command:     "register_service",
		Origin:      origin,
		Destination: "daemon",
		Data:        types.WebsocketSubscription{Service: service},
	}

	return d.Do(request)
}

// GetBlockchainState requests blockchain state
func (d *DaemonService) GetBlockchainState() error {
	request := &types.WebsocketRequest{
		Command:     "get_blockchain_state",
		Origin:      origin,
		Destination: "chia_full_node",
		Data:        map[string]interface{}{},
	}

	return d.Do(request)
}

// GetConnections requests connection info
func (d *DaemonService) GetConnections() error {
	request := &types.WebsocketRequest{
		Command:     "get_connections",
		Origin:      origin,
		Destination: "chia_full_node",
		Data:        map[string]interface{}{},
	}

	return d.Do(request)
}

// GetBlockCountMetrics requests block count metrics
func (d *DaemonService) GetBlockCountMetrics() error {
	request := &types.WebsocketRequest{
		Command:     "get_block_count_metrics",
		Origin:      origin,
		Destination: "chia_full_node",
		Data:        map[string]interface{}{},
	}

	return d.Do(request)
}
