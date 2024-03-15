package address

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"github.com/goethercore/goether/internals"
	"github.com/goethercore/goether/internals/address/address_core"
	"github.com/goethercore/goether/internals/block/block_core"
	"github.com/goethercore/goether/types"
	"github.com/goethercore/goether/utils"
	"github.com/ethereum/go-ethereum/ethclient"

)


func CreateWallet()( error,[]byte){
	 wallet,err:= address_core.GenerateAddress()

	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	return wallet,err
}

func GetBalance(rpc string, walletAddress string) ([]byte, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := address_core.GetAddressBalance(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	response := types.WalletBalanceResponse{
		Balance: walletBalance,
		Address: walletAddress,
	}

	// Marshal the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte
	return responseJSON, nil

}

func GetAddressTransactionCount(rpc string, walletAddress string) (string, error) {
	//"https://bsc.meowrpc.com"
	txCount, err := address_core.GetAddressTXCount(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	return txCount, nil

}

func GetGasPrice(rpc string) (string, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := internals.GasPrice(rpc)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	// Return the JSON-encoded response as a []byte
	return walletBalance, nil

}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := address_core.GetTransactionByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}


	// Marshal the response to JSON
	responseJSON, err := json.Marshal(tx.Result)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionConfirmations(rpc string, hash string) ([]byte, error) {
	tx, err := address_core.GetTransactionByHash(rpc, hash)
	if err != nil {
		return nil, err
	}

	blckNum, err := utils.DecodeBig(tx.Result.BlockNumber)
	if err != nil {
		return nil, err
	}

	block, err := block_core.LastBlock(rpc)
	if err != nil {
		return nil, err
	}

	blockBigInt := new(big.Int)
	_, ok := blockBigInt.SetString(block, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert block number to big.Int")
	}

	height := new(big.Int).Sub(blockBigInt, blckNum)
    value,err:= utils.HexToString(tx.Result.Value)
	if err!=nil{

	}
	confirmation := &types.TransactionConfirmation{
		To:     tx.Result.To,
		From:   tx.Result.From,
		Confirmations: height.String(),
		Amount: value,
	}
	
	confirmationJSON, err := json.Marshal(confirmation)
	if err != nil {
		return nil, err
	}
	
	return confirmationJSON, nil
}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionReceiptByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := address_core.GetTransactionReceiptByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	txData := types.TransactionRecieptResponse{
		JSONRPC: tx.JSONRPC,
		ID:      tx.ID,
		Result:  tx.Result,
	}
	// Marshal the response to JSON
	responseJSON, err := json.Marshal(txData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}
func Transfer(rpc string, privateKey string,recipient string,amount float64, optionalGasPrice ...*big.Int){
	client, err := ethclient.Dial(rpc)
	var gasPrice *big.Int
	if len(optionalGasPrice) > 0 {
		gasPrice = optionalGasPrice[0]
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Failed to suggest gas price: %v", err)
		}
	}
	address_core.SendTx(rpc,privateKey,recipient,amount,gasPrice)

}