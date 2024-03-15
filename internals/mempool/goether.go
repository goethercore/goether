package mempool

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"

	"github.com/goethercore/goether/internals/mempool/mempool_core"
	"github.com/goethercore/goether/types"
)

// TransactionsMemPool retrieves transaction data from the mempool using the provided RPC URL.
func GetMemPoolTransactionsWithStatus(rpc string) ([]byte, error) {
	// Call the mempool.TxMemPool function to fetch the transaction data
	pool, err := mempool_core.MemPoolTransactions(rpc)
	if err != nil {
		// Return an error if there's a problem fetching the transaction data
		return nil, err
	}
	// Create a new MemPoolData struct and populate it with data from the 'pool' variable
	poolData := types.MemPoolTransactionByStatusBlob{
		JSONRPC: pool.JSONRPC,
		ID:      pool.ID,
		Result:  pool.Result,
	}
	// Marshal the transaction data to JSON
	responseJSON, err := json.Marshal(poolData)
	if err != nil {
		// Return an error if there's a problem marshaling to JSON
		return nil, err
	}

	// Return the JSON-encoded transaction data as []byte
	return responseJSON, nil
}


func MemPoolTransactionsCount(rpc string) ([]byte, error) {

	pool, err := mempool_core.MemPoolTransactionsCount(rpc)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	blockData := types.MemPoolTXCount{
		Result: pool.Result,
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
func StreamMempoolTransactions(wssURL string, ch chan string) {

	c,err:= mempool_core.StreamTransactionsMempoolCall(wssURL)
	if err != nil {
		log.Fatal(err)
		
	}
	// Receive and handle JSON-RPC response
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		// Create an EthSubscriptionData object
		var poolData types.MempoolCall

		// Unmarshal the message into the EthSubscriptionData struct
		if err := json.Unmarshal(message, &poolData); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		// Marshal the EthSubscriptionData struct back to JSON
		ethDataJSON, err := json.Marshal(poolData.Params.Result)
		if err != nil {
			log.Println("Error marshaling EthSubscriptionData:", err)
			continue
		}

		// Send the JSON data to the channel as []byte
		ch <- string(ethDataJSON)

	}

	// Handle Ctrl+C to gracefully close the WebSocket connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Closing WebSocket connection...")
}

func ContractTransactionsMempool(wssURL string, contractAddress string, ch chan string) {

	c,err:= mempool_core.StreamTransactionsMempoolCall(wssURL)
	if err != nil {
		log.Fatal(err)
		
	}
	// Receive and handle JSON-RPC response
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		// Create an EthSubscriptionData object
		var poolData types.MempoolCall

		// Unmarshal the message into the EthSubscriptionData struct
		if err := json.Unmarshal(message, &poolData); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		// Marshal the EthSubscriptionData struct back to JSON
		ethDataJSON := poolData.Params.Result
		if err != nil {
			log.Println("Error marshaling EthSubscriptionData:", err)
			continue
		}

		if ethDataJSON.To == strings.ToLower(contractAddress){
			recipientAddress, amount, err := decodeTransactionInput(ethDataJSON.Input)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}// Assign recipient address and amount to the respective fields of ethDataJSON
			ethDataJSON.To = recipientAddress
			ethDataJSON.Value = amount

			ethDataJSONString, err := json.Marshal(ethDataJSON)
			if err != nil {
				log.Println("Error marshaling EthSubscriptionData:", err)
				continue
			}
		
		
		
			ch <- string(ethDataJSONString)

		}
	}

}

func decodeTransactionInput(input string) ( receiver string,amountStr string, err error) {
	if len(input) < 10+64*2 {
		return  "", "",fmt.Errorf("input data length is invalid")
	}

	// Extract receiver address
	receiver = "0x" + input[34:74]

	// Extract amount
	amountBytes, err := hex.DecodeString(input[74:])
	if err != nil {
		return  "","", err
	}
	value := new(big.Int).SetBytes(amountBytes)

    amount := new(big.Float).SetInt(value)

	amountStr = amount.Text('f', 0)

	return receiver,amountStr,  nil
}

