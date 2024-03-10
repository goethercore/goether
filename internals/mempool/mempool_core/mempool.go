package mempool_core

import (
	"encoding/json"

	"net/url"
	"github.com/ayoseun/geth-lte/types"
	"github.com/gorilla/websocket"
)


func StreamTransactionsMempoolCall(wssURL string) (*websocket.Conn, error) {
	// Connect to the WebSocket server
	u, err := url.Parse(wssURL)
	if err != nil {
		return nil, err
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_subscribe",
		Params:  []interface{}{"alchemy_newFullPendingTransactions"},
		ID:      1,
	}

	// Marshal the request to JSON
	requestJSON, err := json.Marshal(request)
	if err != nil {
		c.Close()
		return nil, err
	}

	// Send the JSON-RPC request
	if err := c.WriteMessage(websocket.TextMessage, requestJSON); err != nil {
		c.Close()
		return nil, err
	}

	return c, nil
}
