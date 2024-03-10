package address_core

import (

	"fmt"

	//"fmt"
	
	"github.com/ayoseun/geth-lte/common/hexutil"

)

func GetTransactionConfirmation(rpc string, hash string){
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := GetTransactionByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		

	}



  blckNum,err:=  hexutil.DecodeBig(tx.Result.BlockNumber)

  

fmt.Println(blckNum)
}
