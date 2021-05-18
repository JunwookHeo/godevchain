package main

import (
	"fmt"

	"./block"
)

func main() {
	chain := block.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data in Block: %s\n", block.Data)
	}
}
