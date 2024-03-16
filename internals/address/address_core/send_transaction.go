package address_core

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
	"github.com/goethercore/goether/internals/gas"
	"github.com/goethercore/goether/utils"
)

func SendTx(rpc string,pk string,amount string,  to string) {
    // Replace the following with your private key
    privateKeyHex := pk

    // Recipient address
    toAddress := common.HexToAddress("0xc1b9271024a8512a73481230b94bfbe60e131054")

    // Amount to transfer (0.005 Matic)
    amt,_:= utils.EtherToWei(amount)

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
    noncestr, err := GetAddressTXCount(rpc,fromAddress.String())
    if err != nil {
        log.Fatalf("Failed to retrieve nonce: %v", err)
    }

	nonce, err := strconv.ParseUint(noncestr, 10, 64) // Base 10, 64-bit unsigned integer
    if err != nil {
        fmt.Println("Error:", err)
        return
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
    gasLimit := uint64(21000) // Adjust as needed

    // Create the transaction
    tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amt,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
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
}
