package types

// TransactionResponse represents the JSON response for a transaction.
type TransactionRecieptResponse struct {
    JSONRPC           string             `json:"jsonrpc"`
    ID                int                `json:"id"`
    Result             RecieptResult  `json:"result"`
}

// TransactionResult represents the result data within a transaction response.
type RecieptResult struct {
    BlockHash         string             `json:"blockHash"`
    BlockNumber       string             `json:"blockNumber"`
    ContractAddress   interface{}        `json:"contractAddress"` // It can be a string or null
    CumulativeGasUsed string             `json:"cumulativeGasUsed"`
    EffectiveGasPrice string             `json:"effectiveGasPrice"`
    From              string             `json:"from"`
    GasUsed           string             `json:"gasUsed"`
    Logs              []Log              `json:"logs"`
    LogsBloom         string             `json:"logsBloom"`
    Status            string             `json:"status"`
    To                string             `json:"to"`
    TransactionHash   string             `json:"transactionHash"`
    TransactionIndex  string             `json:"transactionIndex"`
    Type              string             `json:"type"`
}

// Log represents the log data within the logs array.
type Log struct {
    // Define the fields for the log data here
}

