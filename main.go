package main

import (
	"go-blockchain/commandline"
	"os"
)

func main() {
	defer os.Exit(0)

	cli := commandline.CommandLine{}
	cli.Run()

}
