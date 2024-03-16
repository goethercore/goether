package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/goethercore/goether/internals/contract"
	"github.com/goethercore/goether/internals/mempool"
	"github.com/goethercore/goether/types"
	"github.com/goethercore/goether/utils"
)

var rpc = "https://polygon-mumbai.g.alchemy.com/v2/9h8xaWqfXG7B2ENM2fJybKP7glwWM3XH"
var wssrpc = "wss://polygon-mainnet.g.alchemy.com/v2/9h8xaWqfXG7B2ENM2fJybKP7glwWM3XH"
var walletAddress = "0xe9a406f1bb9C0bb1D8Fb8Af3eE50b3C37d1F0Eb2"
var pk = "cb5b800d6310735b8cfd2abc2681cd00ab4b20e4348fd4c1a4b4454df9512172"
var amt = "0.05"
var reciever = "0xC1B9271024a8512A73481230b94bFbe60E131054"
var blockHash = "0xad37be067e06b8c3e2fd741805fac0f82dcad15de5019bd8d0bd2ace73061259"
var hash = "0x0d75b253ef3cdf09c528d4cc47fb5256c1e43d3d9c8d2ddc4c16b65f3cdfcf74"
var contractAddress = "0x202a60A75892CB0EB352fCe2cce5c57EfBFc3CB1"

func main() {
	ListenSmartContractMempoolTx()
}



func ListenSmartContractMempoolTx() {
	poolCh := make(chan string)
    var contractAddress="0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"
	//var rpc = "wss://polygon-mumbai.g.alchemy.com/v2/*************"

	go mempool.ContractMempoolTransactions(wssrpc,contractAddress, poolCh)
	for value := range poolCh {
		var poolData types.MempoolData
		if err := json.Unmarshal([]byte(value), &poolData); err != nil {
			log.Println("error unmarshaling", err)
			continue
		}
		amount, err := utils.HexToString(poolData.Value)

		if err != nil {
			log.Println("error decoding hex value", err)
		}
		fmt.Printf("Transaction From: %s of %s to %s \n", poolData.From, amount, poolData.To)

	}
}


func mutateContract() {

	abi := map[string]string{
		"decimals":      "function decimals()",
		"symbol":        "function symbol()",
		"name":          "function name()",
		"totalSupply":   "function totalSupply()",
		"balanceOf":     "function balanceOf(address)",
		"transfer":      "function transfer(address to, uint256 value)",
		"TransferEvent": "event Transfer(address from, address to, uint256 value)",
	}

	// For functions with argument
	data := map[string]interface{}{
		"functionName": "transfer",
		"args":         []interface{}{reciever, amt}, // corrected syntax
	}

	result, _ := contract.Mutate(rpc, pk, abi, contractAddress, data)

	res, _ := utils.HexToText(result)
	println(res)
}

func readContract() {

	abi := map[string]string{
		"decimals":      "function decimals()",
		"symbol":        "function symbol()",
		"name":          "function name()",
		"totalSupply":   "function totalSupply()",
		"balanceOf":     "function balanceOf(address)",
		"transfer":      "function transfer(address to, uint256 value) external returns (bool)",
		"TransferEvent": "event Transfer(address from, address to, uint256 value)",
	}

	// For functions with argument
	data := map[string]interface{}{
		"functionName": "balanceOf",
		"args":         []interface{}{"0xe9a406f1bb9C0bb1D8Fb8Af3eE50b3C37d1F0Eb2"}, // corrected syntax
	}

	result, _ := contract.Call(rpc, abi, contractAddress, walletAddress, data)
	resultStr, _ := utils.ConvertHexToBigInt(result)

	denominatorStr := "1000000000000000000"
	//setting the precision to 18 is not compulsory, but it defaults to 18
	ethbalance, _ := utils.DivideLargeNumbers(resultStr.String(), denominatorStr, 18)
	println(ethbalance)

}

// func getBlockByHash() {

// 	result,err:=block.GetBlockByHash(rpc,blockHash)

// 		if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("tx Confirmation: %s\n", result)

// }

// func getAllBlockTransactionsCount() {

// 	result,err:=block.GetBlockTransactionCount(rpc,blockHash)

// 		if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("tx Count: %s\n", result)

// }

// func getLatestBlock() {

// 	result,err:=block.GetLatestBlock(rpc)

// 		if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("block: %s\n", result)

// }

// func getTransactionConfirmation() {
// 	result, err := address.GetTransactionConfirmations(rpc, hash)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("tx Confirmation: %s\n", result)
// }

// func getAWalletBalance() {

// 	value, err := address.GetBalance(rpc, walletAddress)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(value)

// }

// func createAddress() {
// 	value, err := address.CreateWallet()
// 	if err != nil {
// 		return
// 	}

// 	var walletData types.Wallet
// 	if err := json.Unmarshal([]byte(value), &walletData); err != nil {
// 		log.Println("error unmarshaling", err)

// 	}

// 	fmt.Println(walletData.PrivateKey)
// }

// func sendCoin() {

// 	value, err := address.Transfer(rpc, pk, amt, reciever)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(value)
// }

// func addressNonce() {
// 	value, err := address.GetAddressTransactionCount(rpc,walletAddress)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println(value)
// }

// func getTransactionByHash() {
// 	result, err := address.GetTransactionByHash(rpc, hash)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var response types.TransactionData
// 	if err :=json.Unmarshal([]byte( result), &response); err !=nil{
// 		log.Println("error unmarshaling", err)
// 	}
// 	fmt.Println(response)
// }
