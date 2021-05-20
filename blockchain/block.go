package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block type
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a block with a hash derived from the data and the prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Genesis returns a genesis block
func Genesis() *Block {
	return CreateBlock("GENESIS", []byte{})
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
