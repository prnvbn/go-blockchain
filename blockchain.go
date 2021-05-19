package main

// BlockChain is a list(chain) of blocks
type BlockChain struct {
	blocks []*Block
}

//AddBlock adds a block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis returns a genesis block
func Genesis() *Block {
	return CreateBlock("GENESIS", []byte{})
}

// InitBlockChain returns a blockchain initialised with the genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
