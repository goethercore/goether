package address_core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/goethercore/goether/types"
)



func GenerateAddress() (error, []byte) {
	// Generate a private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err, nil
	}

	// Get the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err, nil
	}

	// Get the address
	address := publicKeyToAddress(publicKeyECDSA)

	// Get the private key in hex format
	privateKeyBytes := privateKey.D.Bytes()
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	// Create KeyPair struct
	keyPair := types. Wallet{
		Address:    address,
		PublicKey:  fmt.Sprintf("0x%x", elliptic.Marshal(publicKeyECDSA.Curve, publicKeyECDSA.X, publicKeyECDSA.Y)),
		PrivateKey: privateKeyHex,
	}

	// Marshal KeyPair struct to JSON
	keyPairJSON, err := json.Marshal(keyPair)
	if err != nil {
		return err, nil
	}

	return err, keyPairJSON
}

func publicKeyToAddress(pub *ecdsa.PublicKey) string {
	publicKeyBytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	hash := sha256.Sum256(publicKeyBytes[1:]) // Skip the first byte (0x04)
	address := hash[12:]                      // Ethereum address is the last 20 bytes of the hash
	return fmt.Sprintf("0x%x", address)
}
