package block

import (
	"encoding/json"

	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/internals/block/block_core"
	"github.com/ayoseun/geth-lte/types"
)
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
func GetBlockByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	block, err := block_core.GetBlockByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	blockData := types.BlockResponse{
		JSONRPC: block.JSONRPC,
		ID:      block.ID,
		Result:  block.Result,
	}
	// Marshal the response to JSON
	responseJSON, err := json.Marshal(blockData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}
func GetLatestBlock(rpc string) (string, error) {
	//"https://bsc.meowrpc.com"
	block, err := block_core.LastBlock(rpc)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	// Return the JSON-encoded response as a []byte
	return block, nil

}

func GetBlockTransactionCount(rpc string, hash string) (string, error) {
	//"https://bsc.meowrpc.com"
	txCount, err := block_core.GetBlockTXCountByHash(rpc, hash)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	return txCount, nil

}
