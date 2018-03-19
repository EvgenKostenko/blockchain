package main

import (
	"github.com/EvgenKostenko/blockchain/blockchain"
	cliService "github.com/EvgenKostenko/blockchain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Close()

	cli := cliService.New(bc)
	cli.Run()
}
