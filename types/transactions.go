package types


// TransactionResponse represents the JSON response for a transaction.
type TransactionResponse struct {
    JSONRPC string          `json:"jsonrpc"`
    ID      int             `json:"id"`
    Result  TransactionData `json:"result"`
}

// TransactionData represents the data within a transaction response.
type TransactionData struct {
    BlockHash        string `json:"blockHash"`
    BlockNumber      string `json:"blockNumber"`
    From             string `json:"from"`
    Gas              string `json:"gas"`
    GasPrice         string `json:"gasPrice"`
    Hash             string `json:"hash"`
    Input            string `json:"input"`
    Nonce            string `json:"nonce"`
    To               string `json:"to"`
    TransactionIndex string `json:"transactionIndex"`
    Value            string `json:"value"`
    V                string `json:"v"`
    R                string `json:"r"`
    S                string `json:"s"`
}





