package main

import (
	"go-blockchain/blockchain"
	"go-blockchain/commandline"
	"os"
)

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cli := commandline.CommandLine{Blockchain: chain}
	cli.Run()

}
