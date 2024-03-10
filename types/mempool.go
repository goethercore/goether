package types

type MemPoolTXCount struct {
	JSONRPC string             `json:"jsonrpc"`
	ID      int                `json:"id"`
	Result  MemPoolTXCountResult `json:"result"`
}

type MemPoolTXCountResult struct {
	Pending string `json:"pending"`
	Queued  string `json:"queued"`
}


type MemPoolTransactionByStatusData struct {
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	TransactionIndex     interface{}   `json:"transactionIndex"`
	Value                string        `json:"value"`
	Type                 string        `json:"type"`
	AccessList           []interface{} `json:"accessList"`
	ChainID              string        `json:"chainId"`
	V                    string        `json:"v"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	Creates              interface{}   `json:"creates"`
	Wait                 interface{}   `json:"wait"`
}

type ResultData struct {
	Pending map[string]map[string]MemPoolTransactionByStatusData `json:"pending"`
	Queued  map[string]map[string]MemPoolTransactionByStatusData `json:"queued"`
}


type MemPoolTransactionByStatusBlob struct {
	JSONRPC string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  ResultData `json:"result"`
}

type MempoolData struct {
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	To               string  `json:"to"`
	TransactionIndex *string `json:"transactionIndex"`
	Value            string  `json:"value"`
	Type             string  `json:"type"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
}

type MempoolParams struct {
	Result MempoolData `json:"result"`
}

type MempoolCall struct {
	JSONRPC      string        `json:"jsonrpc"`
	Method       string        `json:"method"`
	Params       MempoolParams `json:"params"`
	Subscription string        `json:"subscription"`
}
