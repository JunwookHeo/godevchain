package blockchain

import (
	"container/list"
)

type Chain struct {
	Blocks *list.List
}

func (c *Chain) AddBlock(data string) {
	pb, _ := c.Blocks.Back().Value.(Block)
	new := CreateBlock(data, pb.Hash)
	c.Blocks.PushBack(new)
}

func InitChain() *Chain {
	c := Chain{list.New()}
	c.Blocks.PushBack(Genesis())
	return &c
}
