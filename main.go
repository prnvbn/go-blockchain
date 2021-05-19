package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("FIRST BLOCK AFTER GENESIS")
	chain.AddBlock("SECOND BLOCK AFTER GENESIS")
	chain.AddBlock("THIRD BLOCK AFTER GENESIS")

	for _, b := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", b.PrevHash)
		fmt.Printf("Data         : %s\n", b.Data)
		fmt.Printf("Hash         : %x\n", b.Hash)
	}
}
