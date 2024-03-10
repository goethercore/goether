package interfaces

import (
	"errors"
	"math/big"

	"github.com/ayoseun/geth-lte/common"
	
)

// NotFound is returned by API methods if the requested item does not exist.
var NotFound = errors.New("not found")

// CallMsg contains parameters for contract calls.
type CallMsg struct {
	From      common.Address  // the sender of the 'transaction'
	To        common.Address // the destination contract (nil for contract creation)
	Gas       uint64          // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int        // wei <-> gas exchange ratio
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	Value     *big.Int        // amount of wei sent along with the call
	Data      []byte          // input data, usually an ABI-encoded contract method invocation

	AccessList AccessList // EIP-2930 access list.
}

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"     gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
}