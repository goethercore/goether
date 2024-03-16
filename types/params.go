package types

type ParamObject struct {
	To   string `json:"to"`
	From string `json:"from"`
	Data string `json:"data"`
}

type ParamMutateObject struct {
	Nonce    string `json:"nonce"`
	To       string `json:"to"`
	Value    string `json:"value"`
	GasPrice string `json:"gasPrice"`
	GasLimit string `json:"gasLimit"`
	Data     string `json:"data"`
}
