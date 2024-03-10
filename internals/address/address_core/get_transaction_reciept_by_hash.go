package address_core

import (
	"encoding/json"
	//"fmt"

	"github.com/ayoseun/geth-lte/rpc_calls"
	"github.com/ayoseun/geth-lte/types" // Import the JSONRPC package
)

func GetTransactionReceiptByHash(rpc string, hash string) (types.TransactionRecieptResponse, error) {
    // Define the URL you want to send a POST request to
    url := rpc

    // Create a JSON-RPC request struct
    request := types.JSONRPCRequest{
        JSONRPC: "2.0",
        Method:  "eth_getTransactionReceipt",
        Params:  []interface{}{hash},
        ID:      123,
    }

    // Specify the content type for the request
    contentType := "application/json"

    // Send the JSON-RPC request and handle the response
    response, err := rpccalls.HttpRequest(url, request, contentType)
    if err != nil {
	return types.TransactionRecieptResponse{},nil
    }
	//fmt.Printf("Ether balance: %s\n",response)
    // Define a struct to represent the JSON response
    var parsedResponse types.TransactionRecieptResponse

    // Parse the JSON response into the struct
    err = json.Unmarshal([]byte(response), &parsedResponse)
    if err != nil {
        return types.TransactionRecieptResponse{},nil
    }


    return parsedResponse, nil
}
