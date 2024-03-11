

```markdown
# Goether

`goether` is a Go package designed to provide easy access to Ethereum and Ethereum-compatible blockchain functionalities, such as interacting with smart contracts, reading the mempool for pending transactions, accessing block information, and more.

## Installation

You can install `goether` using `go get`:

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

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This Markdown formatted README provides the same content as the previous template but in Markdown syntax.
