package types

// JSONRPCRequest represents a JSON-RPC request.
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type JSONRPCResult struct {
	ID     int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Error   *RPCError       `json:"error"`
	Result  string `json:"result"`
}

// Ethereum RPC error structure
type RPCError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}