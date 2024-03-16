package gas

import (
	"encoding/json"
	"github.com/goethercore/goether/utils"
	"github.com/goethercore/goether/rpc_calls"
	"github.com/goethercore/goether/types" // Import the JSONRPC package
)

func SuggestGasPrice(rpc string) (string, error) {
	//"https://bsc.meowrpc.com"
	// Define the URL you want to send a POST request to
	url := rpc

// Create a JSON-RPC request struct with an empty array for Params
request := types.JSONRPCRequest{
    JSONRPC: "2.0",
    Method:  "eth_maxPriorityFeePerGas",
    Params:  []interface{}{},
    ID:      123,
}

	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpccalls.HttpRequest(url, request, contentType)
	if err != nil {
		return "", err
	}
	// Create a variable to hold the JSON response
	var parsedResponse types.JSONRPCResult

	// Parse the JSON response into the struct
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {
		return "", err
	}


	resultStr,_ := utils.ConvertHexToBigInt( parsedResponse.Result)
    

	
	return resultStr.String(), nil
}
