package ethcrypto

import (
	"encoding/hex"
	"errors"
	"strings"
	"golang.org/x/crypto/sha3"
)

// calculateFunctionHash calculates the Keccak-256 hash of the given function signature
func calculateFunctionHash(signature string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signature))
	
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateEthereumData generates Ethereum data for calling a function at a specific address
func GenerateEthereumData(args *[]interface{}, functionName string, functionSignatures map[string]string) (string, error) {
	// Find function signature based on function name
	functionSignature, found := functionSignatures[functionName]
	if !found {
		return "", errors.New("function signature not found for the given function name")
	}
	functionSignature = strings.Replace(functionSignature, "function ", "", 1)
	// Calculate function hash (function selector)
	functionHash := calculateFunctionHash(functionSignature)
	data := functionHash[:8] // Take first 4 bytes (8 characters)
	if args != nil && len(*args) > 0 {

		for _, arg := range *args {
			// Convert address to bytes
	
				// Encode the argument based on its type
				encodedArg, err := encodeArgument(arg)
				if err != nil {
					return "", err
				}
				// Concatenate the encoded argument
				data += encodedArg
			

			return "0x" + data, nil
		}

	}

	// Concatenate function signature and padded address


	return "0x" + data, nil
}


func encodeArgument(arg interface{}) (string, error) {
	switch v := arg.(type) {
	case string:
		// If it's a string and starts with "0x", treat it as an Ethereum address
		if strings.HasPrefix(v, "0x") {
			addressBytes, err := hex.DecodeString(v[2:])
			if err != nil {
				return "", err
			}
			paddedAddress := make([]byte, 32)
			copy(paddedAddress[32-len(addressBytes):], addressBytes)
			return hex.EncodeToString(paddedAddress), nil
		}
		// Otherwise, encode it as a regular string
		return hex.EncodeToString([]byte(v)), nil
	case uint64:
		// If it's a uint64, encode it as little-endian
		bytes := make([]byte, 32)
		for i := 0; i < 8; i++ {
			bytes[i] = byte(v >> uint(i*8))
		}
		return hex.EncodeToString(bytes), nil
	case bool:
		// If it's a bool, encode it as 0x01 or 0x00
		if v {
			return "01", nil
		}
		return "00", nil
	default:
		// Unsupported type
		return "", errors.New("unsupported argument type")
	}
}