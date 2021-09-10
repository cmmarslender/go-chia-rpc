# Go Chia RPC

Library for interacting with Chia RPC

## Usage

First, create a new client. Chia config will be automatically read from CHIA_ROOT. If chia is installed under the same user this is running as, it should be automatically discovered.

```go
client, err := rpc.NewClient()
if err != nil {
	// error happened
}
```

Then, just call a method on one of the RPC services

### Get Transactions

```go
transactions, _, err := client.WalletService.GetTransactions(
    &rpc.GetWalletTransactionsOptions{
        WalletID: 1,
    },
)
if err != nil {
    log.Fatal(err)
}

for _, transaction := range transactions.Transactions {
    log.Println(transaction.Name)
}
```

### Get Full Node Status

```go
state, _, err := client.FullNodeService.GetBlockchainState()
if err != nil {
    log.Fatal(err)
}

log.Println(state.BlockchainState.Difficulty)
```

### Get Estimated Network Space

Gets the estimated network space and formats it to a readable version using FormatBytes utility function

```go
//import (
//    "log"
//
//    "github.com/cmmarslender/go-chia-rpc/pkg/rpc"
//    "github.com/cmmarslender/go-chia-rpc/pkg/util"
//)

state, _, err := client.FullNodeService.GetBlockchainState()
if err != nil {
    log.Fatal(err)
}

log.Println(util.FormatBytes(state.BlockchainState.Space))
```

### Request Cache

There is an optional request cache that can be enabled with a configurable cache duration. To use the cache, initialize the client with the `rpc.WithCache()` option like the following example:

```go
client, err := rpc.NewClient(rpc.WithCache(60 * time.Second))
if err != nil {
	// error happened
}
```

This example sets the cache time to 60 seconds. Any identical requests within the 60 seconds will be served from the local cache rather than making another RPC call.
