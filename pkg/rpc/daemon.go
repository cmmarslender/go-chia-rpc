package rpc

import (
	"fmt"
	"github.com/cmmarslender/go-chia-rpc/pkg/types"
	"github.com/gorilla/websocket"
	"net/url"
)

const origin string = "go-chia-rpc"

// DaemonService encapsulates websocket functionality with the daemon
type DaemonService struct {
	client *Client

	conn   *websocket.Conn
}

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

func (d *DaemonService) Do(req *types.WebsocketRequest) error {
	err := d.ensureConnection()
	if err != nil {
		return err
	}

	return d.conn.WriteJSON(req)
}

// Subscribe adds a subscription to a particular service
func (d *DaemonService) Subscribe(service string) error {
	request := &types.WebsocketRequest{
		Command:     "register_service",
		Data:        types.WebsocketSubscription{Service: service},
		Origin:      origin,
		Destination: "daemon",
	}

	return d.Do(request)
}
