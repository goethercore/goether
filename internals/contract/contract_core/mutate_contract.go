package contract_core

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goethercore/goether/eth_crypto"
	"github.com/goethercore/goether/internals/address/address_core"
	"github.com/goethercore/goether/internals/gas"

)

func Mutate(rpc string,pk string,data string,  to string)(string,error) {
    // Replace the following with your private key
    privateKeyHex := pk

    // Recipient address
    toAddress := common.HexToAddress(to)

 

    // Connect to the Polygon (Matic) network
    client, err := ethclient.Dial(rpc)
    if err != nil {
        log.Fatalf("Failed to connect to the Polygon network: %v", err)
    }

    // Create a private key from hex
    privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
    if err != nil {
        log.Fatalf("Failed to create private key: %v", err)
    }

    // Get the public key from the private key
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Failed to convert public key to ECDSA")
    }

    // Get the sender address
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    // Fetch the nonce for the sender address
    noncestr, err := address_core.GetAddressTXCount(rpc,fromAddress.String())
    if err != nil {
        log.Fatalf("Failed to retrieve nonce: %v", err)
    }

	nonce, err := strconv.ParseUint(noncestr, 10, 64) // Base 10, 64-bit unsigned integer
    if err != nil {
        fmt.Println("Error:", err)
        return "",err
    }

    // Set the gas price
    gasPriceStr, err := gas.GasPrice(rpc)
    if err != nil {
        log.Fatalf("Failed to retrieve gas price: %v", err)
    }
    gasPrice, ok := new(big.Int).SetString(gasPriceStr, 10)
    if !ok {
        // Handle error if the string cannot be converted to a big integer
        panic("Failed to convert gasPrice to big integer")
    }
    // Set the gas limit
    gasLimit := uint64(23000) // Adjust as needed

    // Create the transaction
    tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    nil,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:    []byte(data),
		V:        nil,
		R:        nil,
		S:        nil,
	})


    // Sign the transaction
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(80001)), privateKey) // Polygon Mainnet Chain ID is 137
    if err != nil {
        log.Fatalf("Failed to sign transaction: %v", err)
    }

    // Send the transaction
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatalf("Failed to send transaction: %v", err)
    }

    // Output the transaction hash
    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(),err
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
