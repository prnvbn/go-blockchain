package main

import (
	"go-blockchain/wallet"
)

func main() {
	// defer os.Exit(0)

	// cli := commandline.CommandLine{}
	// cli.Run()
	w := wallet.CreateWallet()
	w.Address()

}
