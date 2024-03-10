package types

type WebSocketResponses struct {
    JSONRPC      string `json:"jsonrpc"`
    Method       string `json:"method"`
    Params       Params `json:"params"`
}

type Params struct {
    Result       Result `json:"result"`
    Subscription string `json:"subscription"`
}

type Result struct {
    BlockHash        interface{} `json:"blockHash"`
    BlockNumber      interface{} `json:"blockNumber"`
    From             string      `json:"from"`
    Gas              string      `json:"gas"`
    GasPrice         string      `json:"gasPrice"`
    Hash             string      `json:"hash"`
    Input            string      `json:"input"`
    Nonce            string      `json:"nonce"`
    To               string      `json:"to"`
    TransactionIndex interface{} `json:"transactionIndex"`
    Value            string      `json:"value"`
    V                string      `json:"v"`
    R                string      `json:"r"`
    S                string      `json:"s"`
}
