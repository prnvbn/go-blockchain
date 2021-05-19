package blockchain

// BlockChain is a list(chain) of blocks
type BlockChain struct {
	Blocks []*Block
}

//AddBlock adds a block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// Genesis returns a genesis block
func Genesis() *Block {
	return CreateBlock("GENESIS", []byte{})
}

// InitBlockChain returns a blockchain initialised with the genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
