package main

import (
	"fmt"

	"./block"
)

func main() {
	c := block.InitChain()
	c.AddBlock("First Block after Genesis")
	c.AddBlock("Second Block after Genesis")
	c.AddBlock("Third Block after Genesis")

	for e := c.Blocks.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		b, _ := e.Value.(*block.Block)
		fmt.Printf("Previous Hash: %x\n", b.PreHash)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Data in Block: %s\n", b.Data)
	}

}
