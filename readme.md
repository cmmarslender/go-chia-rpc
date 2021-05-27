# Go Chia RPC

Library for interacting with Chia RPC

:warning: :warning: Go doesn't support anything larger than `uint64`. Many types in Chia are `uint128`. Go will just throw an error once the values get larger than `uint64` can handle :warning: :warning:

# Usage

First, create a new client. Chia config will be automatically read from CHIA_ROOT. If chia is installed under the same user this is running as, it should be automatically discovered.

```go
client, err := rpc.NewClient()
if err != nil {
	// error happened
}
```

Then, just call a method on one of the RPC services

## Get Transactions

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

## Get Full Node Status

```go
state, _, err := client.FullNodeService.GetBlockchainState()
if err != nil {
    log.Fatal(err)
}

log.Println(state.BlockchainState.Difficulty)
```

# Known Issues

* Go doesn't support anything larger than `uint64`. Many types in Chia are `uint128`. Go will just throw an error once the values get larger than `uint64` can handle
