package block

type BlockChain struct {
	Blocks []*Block
}

func (c *BlockChain) AddBlock(data string) {
	pb := c.Blocks[len(c.Blocks)-1]
	new := CreateBlock(data, pb.Hash)
	c.Blocks = append(c.Blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
