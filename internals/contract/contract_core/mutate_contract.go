package contract_core

import (
	//"github.com/ayoseun/geth-lte/internals/address"
	"github.com/ayoseun/geth-lite/rpc_calls"

	"encoding/json"
	//"fmt"

	//"encoding/hex"

	"github.com/ayoseun/geth-lite/types"
)

func SendTransaction(rpc string, signedTx string) (string, error) {
	// Define the URL you want to send a POST request to
	url := rpc

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_sendRawTransaction",
		Params: []interface{}{
			signedTx,
		},
		ID: 1203,
	}
	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpccalls.HttpRequest(url, request, contentType)
	if err != nil {
		return "", err
	}

	// Define a struct to represent the JSON response
	var parsedResponse types.JSONRPCResult

	// Parse the JSON response into the struct
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {
		return "", err
	}

	return parsedResponse.Result, nil
}

// func SendTransactions(rpc string, privateKeyHex string, contractAddressHex string, abiString string, functionName string, params ...interface{}) (string, error) {
// 	// Prepare the function call data
// 	encodedData, err := prepareFunctionCallData(abiString, functionName, params...)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Get the nonce for the sender
// 	pubKey, err := getPublicFromPrivateHelper(privateKeyHex)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonce,err:= address.GetAddressTransactionCount(rpc,pubKey)

// 	// Construct the transaction object
// 	tx := types.Transaction{
// 		Nonce:    nonce,
// 		To:       contractAddressHex,
// 		Value:    "0x0", // Value of Ether being sent, in wei
// 		GasPrice: "0x3B9ACA00", // Example gas price
// 		GasLimit: "0x186A0", // Example gas limit
// 		Data:     encodedData,
// 	}

// 	// Sign the transaction
// 	signedTx, err := signTransaction(tx, privateKeyHex)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Send the signed transaction
// 	txHash, err := sendRawTransaction(rpc, signedTx)
// 	if err != nil {
// 		return "", err
// 	}

// 	return txHash, nil
// }

// func prepareFunctionCallData(abiString string, functionName string, params ...interface{}) (string, error) {
// 	// Encode the function call parameters
// 	encodedParams, err := types.ABIMethodToBytes(abiString, functionName, params...)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Concatenate the function signature and parameters
// 	return hex.EncodeToString(encodedParams), nil
// }

// func getPublicFromPrivateHelper( privateKeyHex string) (string) {
// 	// Get the sender address from the private key
// 	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
// 	if err != nil {
// 		return "", err
// 	}
// 	privateKey, err := crypto.ToECDSA(privateKeyBytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return "", fmt.Errorf("error casting public key to ECDSA")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	return fromAddress.Hex()
// }

// func transactionSigner(tx types.Transaction, privateKeyHex string) (string, error) {
// 	// Sign the transaction
// 	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
// 	if err != nil {
// 		return "", err
// 	}
// 	privateKey, err := crypto.ToECDSA(privateKeyBytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	signedTx, err := types.SignTx(tx, privateKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Marshal the signed transaction
// 	signedTxBytes, err := json.Marshal(signedTx)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(signedTxBytes), nil
// }
