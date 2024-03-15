package mempool_core

import (
	"encoding/json"

	"github.com/goethercore/goether/rpc_calls"
	"github.com/goethercore/goether/types" // Import the JSONRPC package
	"github.com/goethercore/goether/utils"
)

//Gets the total amount of transaction th the mempool
func MemPoolTransactionsCount(rpc string) (types.MemPoolTXCount, error) {
    // Define the URL you want to send a POST request to
    url := rpc

    // Create a JSON-RPC request struct
    request := types.JSONRPCRequest{
        JSONRPC: "2.0",
        Method:  "txpool_status",
        Params:  []interface{}{},
        ID:      123,
    }

    // Specify the content type for the request
    contentType := "application/json"

    // Send the JSON-RPC request and handle the response
    response, err := rpccalls.HttpRequest(url, request, contentType)
    if err != nil {
	
    }

    // Define a struct to represent the JSON response
    var parsedResponse types.MemPoolTXCount



    // Parse the JSON response into the struct
    err = json.Unmarshal([]byte(response), &parsedResponse)
    if err != nil {
     
    }
   // Format the "Pending" and "Queued" values
   parsedResponse.Result.Pending,err = utils.FormatHex(parsedResponse.Result.Pending)
   parsedResponse.Result.Queued,err = utils.FormatHex(parsedResponse.Result.Queued)


    return parsedResponse, nil
}
