package main

import (
	"fmt"
	"strconv"

	"./blockchain"
)

func main() {
	c := blockchain.InitChain()
	for i := 0; i < 10; i++ {
		c.AddBlock(fmt.Sprint("Block : #", i))
	}

	for e := c.Blocks.Front(); e != nil; e = e.Next() {
		b, _ := e.Value.(blockchain.Block)
		fmt.Println("==============================")
		fmt.Printf("Previous Hash: %x\n", b.PrevHash)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Data in Block: %s\n", b.Data)

		pow := blockchain.NewProof(&b)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
