package blockchain

import (
	"go-blockchain/errors"

	"github.com/dgraph-io/badger"
)

// BCIterator is an Iterator for block chain
// TODO? Use interface
type BCIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// Iterator returns a BlockChainIterator for a blockchain
func (chain *BlockChain) Iterator() *BCIterator {
	return &BCIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}
}

// Next will give the next block
func (iter *BCIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		errors.HandleErr(err)
		var encodedBlock []byte
		item.Value(func(val []byte) error {
			encodedBlock = val
			return nil
		})

		block = Deserialize(encodedBlock)

		return nil
	})
	errors.HandleErr(err)

	iter.CurrentHash = block.PrevHash

	return block
}
