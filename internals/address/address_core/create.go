package address_core

import (
	"crypto/ecdsa"
	"encoding/json"
	"log"
    "github.com/goethercore/goether/eth_crypto"
	"github.com/goethercore/goether/types"
	"github.com/goethercore/goether/utils/hexutils"
)



func GenerateAddress() ([]byte,error) {
	privateKey, err := ethcrypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to assert public key to ECDSA type")
	}

	address := ethcrypto.PubkeyToAddress(*publicKeyECDSA)

	privateKeyBytes := privateKey.D.Bytes()  // Convert to byte slice
	 
	privateKeyHex :=hexutils. Encode(privateKeyBytes)

	addressHex :=hexutils. Encode(address[:])
  
	// Create KeyPair struct
	keyPair := types. Wallet{
		Address:    addressHex,
		PrivateKey: privateKeyHex,
	}

	// Marshal KeyPair struct to JSON
	keyPairJSON, err := json.Marshal(keyPair)
	if err != nil {
		return nil,err
	}

	return  keyPairJSON,err
}

