# go-blockchain

A commandline interface for a simple blockchain module in Go. This project is an academic pursuit of mine to learn more about blockchain and implement one on myself. The [charts](charts) folder has some great diagrams that can be used to understand certain concepts of blockchain. 

This module uses SHA256 hashing and so obviously should not be used for sensitive data (SHA256 can be cracked with relative ease). The module uses [BadgerDB](https://github.com/dgraph-io/badger) to store the blockchain and ensure persistance. 

### Transactions

As of now, 1 block can only have 1 transaction because the mining of the block is synchronous with the sending of a transaction. Each block has an input and an output. The input stores an unhashed public key and a signature of the entire transaction and the output stores a hashed public key. So to access the unspent tokens in the output, the hash of the public key in the input has to match the hashed public key in the output. You can see the flowchart of an example transaction [here](charts/transactions.png)

## CLI
There is a simple commandline application showing the module can be used. You can run it using `go run main.go (flags)`. See Usage to learn about the flags.

### Usage
1. `printchain` Prints all the blocks in the chain
2. `getbalance -address ADDRESS` gets the balance for a given address
3. `createblockchain -address ADDRESS` creates a blockchain
4. `send -from FROM -to TO -amount -AMOUNT` makes a transaction
5. `createwallet` - Creates a new Wallet
6. `listaddresses` - Lists the addresses in our wallet file

## Demo
I am assuming you have go properly installed on your machine.

### With wallets

Creating 2 wallets
![Screenshot from 2021-06-13 16-39-21](https://user-images.githubusercontent.com/55818107/121814190-a19c3980-cc67-11eb-8463-131a91ac8727.png)

Mining the genesis block
![Screenshot from 2021-06-13 16-39-41](https://user-images.githubusercontent.com/55818107/121814202-acef6500-cc67-11eb-9e0c-c39208bfdaee.png)

Making a transaction
![Screenshot from 2021-06-13 16-55-13](https://user-images.githubusercontent.com/55818107/121814304-2ab37080-cc68-11eb-9d1e-1b9ba245f53a.png)

### Without wallets
First, we created the blockchain by proving the genesis block and subsequently having a coinbase transaction mining it. After this step, the balance at address "Pranav" is 100 (value of the coinbase transaction). Then we make a transaction from address "Pranav" to address "Manya". This first created an address "Manya" and then completed the transaction. We can see that it was succesful and that the blockchain is still valid after it.

![image](https://user-images.githubusercontent.com/55818107/119258318-11675900-bbc1-11eb-87db-e147c3c1cacf.png)

If you find any issues, typos, errors or have any feature requests, feel free to create an issue or a pull request :)

## TODO
- [x] Persistance
- [X] Transactions
- [ ] Consensus Algorithm
- [X] Wallet Module
- [ ] Connect wallet Module and the blockchain (In Progress)
- [ ] Digital Signatures
- [ ] Merkle Tree
- [ ] Dynamic Difficulty
- [ ] Improve the CLI Package

Being added with wallets. 
![image](https://user-images.githubusercontent.com/55818107/119389459-9175e780-bcc3-11eb-97e2-9d6c87903f80.png)


