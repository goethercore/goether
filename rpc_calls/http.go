package rpccalls

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/ayoseun/geth-lte/types"
)

// GetBalanceRequest sends a JSON-RPC request to the specified URL
// with the given JSON request and content type, and returns the response body or an error.
func HttpRequest(url string, request types .JSONRPCRequest, contentType string) ([]byte, error) {
	// Serialize the JSON-RPC request
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Send the POST request with the JSON-RPC request in the body
	response, err := http.Post(url, contentType, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()



	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
