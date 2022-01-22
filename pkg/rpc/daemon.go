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
	client *Client
	conn   *websocket.Conn
	handler websocketRespHandler
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

// Do sends a request over the websocket connection
func (d *DaemonService) Do(req *types.WebsocketRequest) error {
	err := d.ensureConnection()
	if err != nil {
		return err
	}

	return d.conn.WriteJSON(req)
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
