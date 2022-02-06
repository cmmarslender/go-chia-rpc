package rpcinterface

import "github.com/cmmarslender/go-chia-rpc/pkg/types"

// WebsocketResponseHandler is a function that is called to process a received websocket response
type WebsocketResponseHandler func(*types.WebsocketResponse, error)
