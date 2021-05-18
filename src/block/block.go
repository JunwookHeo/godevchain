package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash    []byte
	PreHash []byte
	Data    []byte
}

func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.PreHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, preHash []byte) *Block {
	block := &Block{[]byte{}, preHash, []byte(data)}
	block.deriveHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
