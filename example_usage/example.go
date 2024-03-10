package exampleusage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/ayoseun/geth-lte/common/hexutil"
	"github.com/ayoseun/geth-lte/internals/address"
	"github.com/ayoseun/geth-lte/internals/block"
	"github.com/ayoseun/geth-lte/internals/contract"
	"github.com/ayoseun/geth-lte/internals/mempool"
	"github.com/ayoseun/geth-lte/types"
	"github.com/ayoseun/geth-lte/utils"
	"github.com/joho/godotenv"
)

var (
	contractCoinAddress2 = "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"
	hash                 = "0x9e4cc336022fd3fdae0f5ad25b758f040f30040a73a45fca1be9e440bac91902"
	wallet               = "0xa6f79B60359f141df90A0C745125B131cAAfFD12"
	DAIContract          = "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"
	addressTestNet       = "0xD48a3323E0349912185Ae4522F083bcc011CEa07"
	rpc                  string
	wssRPC               string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	apiKey := os.Getenv("APIKEY")
	rpc = fmt.Sprintf("https://polygon-mumbai.g.alchemy.com/v2/%s", apiKey)
	wssRPC = fmt.Sprintf("wss://polygon-mumbai.g.alchemy.com/v2/%s", apiKey)

}

func ContractmemPool() {
	poolCh := make(chan string)
	go mempool.ContractTransactionsMempool(wssRPC, "0x202a60A75892CB0EB352fCe2cce5c57EfBFc3CB1", poolCh)
	for value := range poolCh {

		var poolData types.MempoolData
		if err := json.Unmarshal([]byte(value), &poolData); err != nil {
		}
		amount,_ := hexutil.ConvertToEtherDecimal(poolData.Value, 18)

		log.Printf("Found a transaction: Amount: %s from: %s to %s\n", amount, poolData.From, poolData.To)
	}
}

func StreamMemPool() {
	poolCh := make(chan string)
	go mempool.StreamMempoolTransactions(wssRPC, poolCh)
	for value := range poolCh {
		var poolData types.MempoolData
		if err := json.Unmarshal([]byte(value), &poolData); err != nil {
		}
		amount, err := hexutil.DecodeBig(poolData.Value)
		if err != nil {
			continue
		}

		log.Printf("Found a transaction: Amount: %s from: %s to %s\n", amount, poolData.From, poolData.To)

	}
}

func UserContract() {

	abi := map[string]string{
		"decimals":      "function decimals()",
		"symbol":        "function symbol() view returns (string)",
		"name":          "function name()",
		"totalSupply":   "function totalSupply()",
		"balanceOf":     "function balanceOf(address)",
		"get":           "function Get()",
		"transfer":      "function transfer(address to, uint256 value) external returns (bool)",
		"TransferEvent": "event Transfer(address from, address to, uint256 value)",
	}
	// For functions without an argument
	data := map[string]interface{}{
		"functionName": "get",
		"args":         []interface{}{}, // corrected syntax
	}

	// For functions with argument
	// data2 := map[string]interface{}{
	// 	"functionName": "get",
	// 	"args":         []interface{}{wallet,20}, // corrected syntax
	// }

	result, err := contract.Call(rpc, abi, DAIContract, wallet, data)

	if err != nil {
		panic(err)
	}
	res, err := utils.HexToText(result)
	println(res)
}

// Send Native Tokens
func SendCoin() {

	privateKey := ""
	receipent := wallet

	address.Transfer(rpc, privateKey, receipent, 0.056)

}

// Reading mempool for transactions and get quequed and pending transactions
func MemPoolWithStatus() {
	result, err := mempool.GetMemPoolTransactionsWithStatus(rpc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pool: %s\n", result)
}

// Get block hash
func BlockByHash() {
	result, err := block.GetBlockByHash(rpc, "0x499d2f7bcd2c37e869f6721edb690105d19275e2ae25911c7d81b75305075dcd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// Get Transaction hash
func GetTransactionByHash() {
	result, err := address.GetTransactionByHash(rpc, hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// Get Transaction hash
func GetTransactionConfirmation() {
	result, err := address.GetTransactionConfirmations(rpc, hash)

	if err != nil {
		panic(err)
	}
	fmt.Printf("tx Confirmation: %s\n", result)
}

// Get block transaction count
func GetBlockTransactionCounts() {
	result, err := address.GetAddressTransactionCount(rpc, hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// latest block example
func LatestBlock() {
	result, err := block.GetLatestBlock(rpc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Latest Block: %s\n", result)
}

// Get address transactions count
func AddressTransactionCount() {
	result, err := address.GetAddressTransactionCount(rpc, wallet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// Get Wallet Balance
func GetWalletBalance() {
	result, err := address.GetBalance(rpc, wallet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance: %s\n", result)

}
