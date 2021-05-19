package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block type
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash derives the hash for the given block using prevHash and the block data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) //TODO: Replace sha256
	b.Hash = hash[:]
}

// CreateBlock creates a block with a hash derived from the data and the prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}
