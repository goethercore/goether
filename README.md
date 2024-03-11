


# Goether [![Go Reference](https://pkg.go.dev/badge/github.com/ayoseun/geth-lte@v0.0.0-20240207011905-a59fb5d8d2f9.svg)](https://pkg.go.dev/github.com/ayoseun/geth-lte@v0.0.0-20240207011905-a59fb5d8d2f9)

goether is a Go package designed to provide various functionalities related to EVM blockchains. This package offers capabilities such as interacting with contracts, exploring blocks, monitoring mempool transactions in real-time, and more. Whether you're building decentralized applications or exploring blockchain data, geth-lte aims to simplify your development process.

Don't forget to star ⭐️ the repo

## Features

1. Contract Interaction
View Contract Transaction Mempool: Monitor real-time transactions in the mempool associated with a specific smart contract.
2. Blockchain Exploration
Get Block by Hash: Retrieve detailed information about a block using its hash.
Latest Block: Obtain data regarding the latest block on the blockchain.
3. Transaction Handling
Get Transaction by Hash: Retrieve transaction details using its hash.
Transaction Confirmations: Check the number of confirmations for a specific transaction.
Address Transaction Count: Get the total number of transactions associated with a particular address.
4. Wallet Operations
Get Wallet Balance: Retrieve the balance of a wallet address.
5. Mempool Monitoring
View Transaction Mempool: Observe real-time transactions in the mempool of the blockchain network.
Getting Started

## Installation

You can install `goether` using the go installation command `go get`:

```shell
go get github.com/goethercore/goether

```

## Usage

### Initialization

```go
import "github.com/goethercore/goether"

func main() {
    goether.Init()
}
```

### Examples

#### 1. Reading Contract Transactions from Mempool

```go
goether.ContractmemPool()
```

#### 2. Streaming Mempool Transactions

```go
goether.StreamMemPool()
```

#### 3. Retrieving User Contract Information

```go
goether.UserContract()
```

#### 4. Sending Native Tokens

```go
goether.SendCoin()
```

#### 5. Reading Mempool with Status

```go
goether.MemPoolWithStatus()
```

#### 6. Retrieving Block Information

```go
goether.BlockByHash()
```

#### 7. Retrieving Transaction Information

```go
goether.GetTransactionByHash()
```

#### 8. Retrieving Transaction Confirmation

```go
goether.GetTransactionConfirmation()
```

#### 9. Retrieving Block Transaction Counts

```go
goether.GetBlockTransactionCounts()
```

#### 10. Retrieving Latest Block

```go
goether.LatestBlock()
```

#### 11. Retrieving Address Transactions Count

```go
goether.AddressTransactionCount()
```

#### 12. Retrieving Wallet Balance

```go
goether.GetWalletBalance()
```

## Configuration

Before using the package, make sure to set up your environment variables by creating a `.env` file and providing the required API key:

```sh
APIKEY=your-api-key
```

## Contributing

Contributions to geth-lte are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on GitHub.

- Active contributors
  - [Ayoseun](github.com/ayoseun)


## License

This project is licensed under the MIT License - see the LICENSE file for details.
