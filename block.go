package main

import (
	"bytes"
	"crypto/sha256"
)

// Block type
type Block struct {
	hash     []byte
	data     []byte
	prevHash []byte
}

// DeriveHash derives the hash for the given block using prevHash and the block data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.data, b.prevHash}, []byte{})
	hash := sha256.Sum256(info) //TODO: Replace sha256
	b.hash = hash[:]
}

// CreateBlock creates a block with a hash derived from the data and the prevHash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		data:     []byte(data),
		prevHash: prevHash,
	}
	block.DeriveHash()
	return block
}
