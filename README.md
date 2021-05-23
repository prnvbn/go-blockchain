# go-blockchain

A simple blockchain module in Go. 

This module uses SHA256 hashing and so obviously should not be used for sensitive data (SHA256 can be cracked with relative ease). The module uses [BadgerDB](https://github.com/dgraph-io/badger) to store the blockchain and ensure persistance. 

## CLI
There is a simple commandline application showing the module can be used. You can run it using `go run main.go (flags)`. See Usage to learn about the flags.

### Usage
1. `printchain` Prints all the blocks in the chain
2. `getbalance -address ADDRESS` gets the balance for a given address
3. `createblockchain -address ADDRESS` creates a blockchain
4. `send -from FROM -to TO -amount -AMOUNT` makes a transaction



## TODO
- [x] Persistance (BadgerDB)
- [X] Transactions
- [ ] Consensus Algorithm
- [ ] Wallet Module (In Progress)
- [ ] Digital Signatures
- [ ] Merkle Tree

If you find any issues or have an feature requests feel free to create an issue or a pull request :)
