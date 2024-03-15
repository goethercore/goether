package address_core

import (
	"fmt"

	"github.com/goethercore/goether/utils"
)

func GetTransactionConfirmation(rpc string, hash string){
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := GetTransactionByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		

	}



  blckNum,err:=  utils.DecodeBig(tx.Result.BlockNumber)

  

fmt.Println(blckNum)
}
