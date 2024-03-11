package block_core

import (
	"encoding/json"
	//"fmt"

	"github.com/goethercore/goether/rpc_calls"
	"github.com/goethercore/goether/types" // Import the JSONRPC package
)

func GetBlockByHash(rpc string, hash string) (types.BlockResponse, error) {
    // Define the URL you want to send a POST request to
    url := rpc

    // Create a JSON-RPC request struct
    request := types.JSONRPCRequest{
        JSONRPC: "2.0",
        Method:  "eth_getBlockByHash",
        Params:  []interface{}{hash,false},
        ID:      123,
    }

    // Specify the content type for the request
    contentType := "application/json"

    // Send the JSON-RPC request and handle the response
    response, err := rpccalls.HttpRequest(url, request, contentType)
    if err != nil {
		return types.BlockResponse{},nil
    }
	//fmt.Printf("Ether balance: %s\n",response)
    // Define a struct to represent the JSON response
    var parsedResponse types.BlockResponse

    // Parse the JSON response into the struct
    err = json.Unmarshal([]byte(response), &parsedResponse)
    if err != nil {
       return types.BlockResponse{},nil
    }


    return parsedResponse, nil
}
