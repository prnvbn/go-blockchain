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

## Demo
I am assuming you have go properly installed on your machine. 

First, we created the blockchain by proving the genesis block and subsequently having a coinbase transaction mining it. After this step, the balance at address "Pranav" is 100 (value of the coinbase transaction). Then we make a transaction from address "Pranav" to address "Manya". This first created an address "Manya" and then completed the transaction. We can see that it was succesful and that the blockchain is still valid after it.

![image](https://user-images.githubusercontent.com/55818107/119258318-11675900-bbc1-11eb-87db-e147c3c1cacf.png)

## TODO
- [x] Persistance
- [X] Transactions
- [ ] Consensus Algorithm
- [ ] Wallet Module (In Progress)
- [ ] Digital Signatures
- [ ] Merkle Tree
- [ ] Dynamic Difficulty
If you find any issues, typos, errors or have any feature requests, feel free to create an issue or a pull request :)
