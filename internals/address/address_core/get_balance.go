package address_core

import (
	"encoding/json"
	"github.com/goethercore/goether/rpc_calls"
	"github.com/goethercore/goether/types" // Import the JSONRPC package
	"github.com/goethercore/goether/utils"
)

func GetAddressBalance(rpc string, address string) (string, error) {
	//"https://bsc.meowrpc.com"
	// Define the URL you want to send a POST request to
	url := rpc

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBalance",
		Params:  []interface{}{address, "latest"},
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

	denominatorStr := "1000000000000000000"
	//setting the precision to 18 is not compulsory, but it defaults to 18 
	ethbalance, err := utils.DivideLargeNumbers(resultStr.String(), denominatorStr,18)
	if err != nil {
		return "", err
	}
	
	return ethbalance, nil
}
