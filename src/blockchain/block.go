package blockchain

type Block struct {
	Hash     []byte
	Nonce    int
	PrevHash []byte
	Data     []byte
}

// func (b *Block) deriveHash() {
// 	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) Block {
	block := Block{[]byte{}, 0, prevHash, []byte(data)}
	pow := NewProof(&block)
	n, h := pow.Run()
	block.Hash = h[:]
	block.Nonce = n

	return block
}

func Genesis() Block {
	return CreateBlock("Genesis", []byte{})
}
