# Go Chia RPC

Library for interacting with Chia RPC

# Known Issues

* Go doesn't support anything larger than uint64. Many types in Chia are uint128. Go will just throw an error once the values get larger than a uint64 can handle
