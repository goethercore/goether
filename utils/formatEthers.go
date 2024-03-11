package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	hexutil "github.com/goethercore/goether/common/hexutil"
)

func FormatEther(hex string,precision ...int)(string, error) {

	result, err := hexutil.DecodeBig(hex)
	if err != nil {
		return "", err
	}

	denominatorStr := "1000000000000000000"

	resultStr := result.String()
		// Default precision to 8 if not provided
		if len(precision) == 0 {
			precision = append(precision, 18)
		}
	etherStr, err := hexutil.DivideLargeNumbers(resultStr, denominatorStr,precision[0])
	if err != nil {
		return "", err
	}

	return etherStr, nil
}


func ParseEther(big string,precision ...int)(string, error) {
	// Default precision to 8 if not provided
	if len(precision) == 0 {
		precision = append(precision, 18)
	}
	
	etherStr, err := hexutil.MultiplyDecimal(big,precision[0])
	if err != nil {
		return "", err
	}

	return etherStr, nil
}

func HexToDecimal(hexStr string) (*big.Int, error) {
	// Remove the "0x" prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Convert hex string to a big.Int
	decimalValue := new(big.Int)
	decimalValue, success := decimalValue.SetString(hexStr, 16)
	if !success {
		return nil, fmt.Errorf("Invalid hexadecimal string: %s", hexStr)
	}

	return decimalValue, nil
}
func HexToString(hexStr string,precision ...int) (string, error) {
	// Remove the "0x" prefix if present
    decimalValue, err := HexToDecimal(hexStr)
    if err !=nil {
		return "", fmt.Errorf("Invalid hexadecimal string: %s", hexStr)
	}
    denominatorStr := "1000000000000000000"
		// Default precision to 8 if not provided
		if len(precision) == 0 {
			precision = append(precision, 18)
		}
    result,err:= hexutil.DivideLargeNumbers( decimalValue.String(),denominatorStr,precision[0])

    

	return result, nil
}

func HexToText(hexString string) (string, error) {
	// Remove "0x" prefix from the hexadecimal string
	if strings.HasPrefix(hexString, "0x") {
		hexString = hexString[2:]
	}

	// Convert hexadecimal string to bytes
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", fmt.Errorf("error decoding hexadecimal string: %v", err)
	}

	text := string(bytes)

	return text, nil
}