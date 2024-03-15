package utils

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

)


func FormatEther(hex string,precision ...int)(string, error) {

	result, err := DecodeBig(hex)
	if err != nil {
		return "", err
	}

	denominatorStr := "1000000000000000000"

	resultStr := result.String()
		// Default precision to 8 if not provided
		if len(precision) == 0 {
			precision = append(precision, 18)
		}
	etherStr, err := DivideLargeNumbers(resultStr, denominatorStr,precision[0])
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
	
	etherStr, err := MultiplyDecimal(big,precision[0])
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
    result,err:= DivideLargeNumbers( decimalValue.String(),denominatorStr,precision[0])

    

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

func ConvertToEtherDecimal(amount string, decimal int) (string, error) {
	
	divisor := math.Pow10(decimal)
	num, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}
	converted := num / divisor
	return fmt.Sprintf("%.32f", converted), nil
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}


func checkNumber(input string) (raw string, err error) {
	if len(input) == 0 {
		return "", ErrEmptyString
	}
	if !has0xPrefix(input) {
		return "", ErrMissingPrefix
	}
	input = input[2:]
	if len(input) == 0 {
		return "", ErrEmptyNumber
	}
	if len(input) > 1 && input[0] == '0' {
		return "", ErrLeadingZero
	}
	return input, nil
}
type decError struct{ msg string }

func (err decError) Error() string { return err.msg }


func decodeNibble(in byte) uint64 {
	switch {
	case in >= '0' && in <= '9':
		return uint64(in - '0')
	case in >= 'A' && in <= 'F':
		return uint64(in - 'A' + 10)
	case in >= 'a' && in <= 'f':
		return uint64(in - 'a' + 10)
	default:
		return badNibble
	}
}

// DecodeBig decodes a hex string with 0x prefix as a quantity.
// Numbers larger than 256 bits are not accepted.
func DecodeBig(input string) (*big.Int, error) {
	raw, err := checkNumber(input)
	if err != nil {
		return nil, err
	}
	if len(raw) > 64 {
		return nil, ErrBig256Range
	}
	words := make([]big.Word, len(raw)/bigWordNibbles+1)
	end := len(raw)
	for i := range words {
		start := end - bigWordNibbles
		if start < 0 {
			start = 0
		}
		for ri := start; ri < end; ri++ {
			nib := decodeNibble(raw[ri])
			if nib == badNibble {
				return nil, ErrSyntax
			}
			words[i] *= 16
			words[i] += big.Word(nib)
		}
		end = start
	}
	dec := new(big.Int).SetBits(words)
	return dec, nil
}

func DivideLargeNumbers(numeratorStr, denominatorStr string, precision int) (string, error) {
	// Create big.Int instances representing the numerator and denominator
	num := new(big.Int)
	denom := new(big.Int)

	// Set the values of the numerator and denominator
	if _, success := num.SetString(numeratorStr, 10); !success {
		return "", fmt.Errorf("invalid numerator: %s", numeratorStr)
	}
	if _, success := denom.SetString(denominatorStr, 10); !success {
		return "", fmt.Errorf("invalid denominator: %s", denominatorStr)
	}



	// Perform the division
	var result big.Float
	result.Quo(new(big.Float).SetInt(num), new(big.Float).SetInt(denom))

	// Convert the result to a string with the specified precision
	resultStr := result.Text('f', precision)

	return resultStr, nil
}

func MultiplyDecimal(denominatorStr string, precision int) (string, error) {
    // Convert the input string to a big.Float
    denominator, _, err := new(big.Float).Parse(denominatorStr, 10)
    if err != nil {
        return "", err
    }

    // Create a big.Float representing 10 raised to the power of the precision
    tenToThePower := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil))

    // Multiply the denominator by 10^precision
    result := new(big.Float).Mul(denominator, tenToThePower)

    // Convert the result to a string with the specified precision
    resultStr := result.Text('f', precision)

    return resultStr, nil
}

func FormatHex(hex string)(string, error) {

	result, err := DecodeBig(hex)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}