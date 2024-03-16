package address

import (

	"encoding/json"
	"fmt"
	"github.com/goethercore/goether/internals/address/address_core"
	"github.com/goethercore/goether/internals/block/block_core"
	"github.com/goethercore/goether/types"
	"github.com/goethercore/goether/utils"

	"math/big"
)

func CreateWallet() ([]byte,error) {
	wallet, err := address_core.GenerateAddress()

	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil,err
	}

	return  wallet,err
}

func GetBalance(rpc string, walletAddress string) (string, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := address_core.GetAddressBalance(rpc, walletAddress)

	// Return the JSON-encoded response as a []byte
	return walletBalance,err

}

func GetAddressTransactionCount(rpc string, walletAddress string) (string, error) {

	txCount, err := address_core.GetAddressTXCount(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	return txCount, nil

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

	blckNum, err := utils. ConvertHexToBigInt(tx.Result.BlockNumber)
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
	value, err := utils.HexToString(tx.Result.Value)
	if err != nil {

	}
	confirmation := &types.TransactionConfirmation{
		To:            tx.Result.To,
		From:          tx.Result.From,
		Confirmations: height.String(),
		Amount:        value,
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
func Transfer(rpc string, privateKey string,  amount string,recipient string, optionalGasPrice ...*big.Int)(string,error) {

address_core.SendTx(rpc,privateKey,amount,recipient)
	return "",nil

}
