package contract

import (
	"fmt"
	"github.com/goethercore/goether/eth_crypto"
	"github.com/goethercore/goether/types"
	"github.com/goethercore/goether/internals/contract/contract_core"

)

func Read(rpc string, abi map[string]string, contractAddress string,address string, data map[string]interface{}) (string, error) {
	// Define the RPC string

	functionName := data["functionName"].(string)
	args := data["args"].([]interface{})

	ethereumData, err := ethcrypto.GenerateEthereumData(&args, functionName, abi)
	if err != nil {
		fmt.Println("Error:", err)

	}
	//fmt.Println("Ethereum Data:", ethereumData)

	// Create an instance of CallMsg and fill it with data
	msg := types.ParamObject{
		To:   contractAddress,
		From: address,
		Data: ethereumData,
	}
	// Call the BalanceOf function
	resp, err := contract_core.Call(rpc, msg)
	if err != nil {

	}

	

	return resp, nil

}

func Write(rpcURL string,privateKey string,  ABI map[string]string, contractAddress string, data map[string]interface{}) (string, error) {
	// Define the RPC string

	functionName := data["functionName"].(string)
	args := data["args"].([]interface{})

	ethereumData, err := ethcrypto.GenerateEthereumData(&args, functionName, ABI)
	if err != nil {
		fmt.Println("Error:", err)

	}

	// Call the BalanceOf function
	resp, err := contract_core.Mutate(rpcURL,privateKey,ethereumData,contractAddress)
	if err != nil {

	}

	

	return resp, nil

}