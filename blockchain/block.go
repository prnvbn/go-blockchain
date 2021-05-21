package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Block type
type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

// CreateBlock creates a block with a hash derived from the data and the prevHash
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
	}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Genesis returns a genesis block
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// HashTransactions returns the hash of all the transactions in the block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// Serialize serializes the block into a byte slice
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	HandleErr(err)

	return res.Bytes()
}

// Deserialize deserializes the byte slice into a block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	HandleErr(err)
	return &block
}

// HandleErr panics and logs the error
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
