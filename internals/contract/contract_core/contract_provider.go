package contract_core


import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTokenTx(rpc string, pk string, recipient string, amount *big.Int, contractAddress string, contractABI string, optionalGasPrice ...*big.Int) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(string(pk))
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	toAddress := common.HexToAddress(string(recipient))
	data, err := parsedABI.Pack("transfer", toAddress, amount)
	if err != nil {
		log.Fatalf("Failed to pack data for transfer: %v", err)
	}

	gasLimit := uint64(200000) // set a higher gas limit for token transfers
	var gasPrice *big.Int
	if len(optionalGasPrice) > 0 {
		gasPrice = optionalGasPrice[0]
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Failed to suggest gas price: %v", err)
		}
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign tx: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	log.Printf("Tx sent: %s", signedTx.Hash().Hex())
}