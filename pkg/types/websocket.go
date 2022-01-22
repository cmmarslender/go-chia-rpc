package types

// WebsocketRequest represents the common elements of a websocket request and response
type WebsocketRequest struct {
	Command     string      `json:"command"`
	Data        interface{} `json:"data"`
	Ack         bool        `json:"ack"`
	Origin      string      `json:"origin"`
	Destination string      `json:"destination"`
	RequestID   string      `json:"request_id"`
}

// WebsocketSubscription is the Data for a new subscribe request
type WebsocketSubscription struct {
	Service string `json:"service"`
}
