package contract_core

import (
	"encoding/json"
	 "github.com/ayoseun/geth-lite/rpc_calls"

	"github.com/ayoseun/geth-lite/types" // Import the JSONRPC package
)

func Call(rpc string, msg types.ParamObject) (string, error) {
	// Define the URL you want to send a POST request to
	url := rpc

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  ToCallArg(msg),

		ID: 1203,
	}

	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpccalls.HttpRequest(url, request, contentType)
	if err != nil {

	}

	// Define a struct to represent the JSON response
	var parsedResponse types.JSONRPCResult

	// Parse the JSON response into the struct
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {

	}

	

	return parsedResponse.Result, nil
}
