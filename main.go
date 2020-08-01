package main

import (
	"fmt"
	"strconv"
	"github.com/AbDwivedi7/block-chain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n",block.PrevHash)
		fmt.Printf("Data in the block: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("\n")
	}
	
}