package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// BlockChain is a list(chain) of blocks
type BlockChain struct {
	// Blocks   []*Block
	LastHash []byte //Last hash of the last block in the chain
	Database *badger.DB
}

// InitBlockChain returns a blockchain initialised with the genesis block
func InitBlockChain() *BlockChain {

	// opening the database
	opts := badger.DefaultOptions(dbPath)
	opts.Logger = nil
	db, err := badger.Open(opts)
	HandleErr(err)

	var lastHash []byte
	err = db.Update(func(txn *badger.Txn) error {
		// check if blockchain exists
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			// It does not!

			// Creating Genesis block
			fmt.Println("No existing blockchain")
			genesis := Genesis()
			fmt.Println("Genesis Proved")

			// Adding the Genesis block
			err = txn.Set(genesis.Hash, genesis.Serialize())
			HandleErr(err)

			// Updating the last hash
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash
			return err
		} else {
			// It does!

			// Getting the last hash
			item, err := txn.Get([]byte("lh"))
			HandleErr(err)
			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})
	HandleErr(err)

	return &BlockChain{
		Database: db,
		LastHash: lastHash,
	}
}

//AddBlock adds a block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	// prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// newBlock := CreateBlock(data, prevBlock.Hash)
	// chain.Blocks = append(chain.Blocks, newBlock)
	var lastHash []byte

	// Getting the last hash and creating a new block
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleErr(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})
	HandleErr(err)
	newBlock := CreateBlock(data, lastHash)

	// Updating the last hash key
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		HandleErr(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		chain.LastHash = newBlock.Hash
		return err
	})
}
