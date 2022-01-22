package rpc

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

// DaemonService encapsulates websocket functionality with the daemon
type DaemonService struct {
	client *Client

	conn   *websocket.Conn
}

func (d *DaemonService) ensureConnection() error {
	if d.conn == nil {
		u := url.URL{Scheme: "wss", Host: fmt.Sprintf("%s:%d", d.client.baseURL.Host, d.client.daemonPort), Path: "/"}
		var err error
		d.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DaemonService) Do() error {
	err := d.ensureConnection()
	if err != nil {
		return err
	}

	// Do!

	return nil
}
