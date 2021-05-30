package blockchain

import (
	"crypto/sha256"
)

func CalHashSha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func CalMerkleNodeHash(l, r []byte) []byte {
	return CalHashSha256(append(l, r...))
}

func CalMerkleUpperHashs(hashes [][]byte) [][]byte {
	var mtns [][]byte

	if len(hashes)%2 == 1 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	for i := 0; i < len(hashes); i += 2 {
		node := CalMerkleNodeHash(hashes[i], hashes[i+1])
		mtns = append(mtns, node)
	}

	return mtns
}

func CalMerklRootHash(data [][]byte) []byte {
	var mtns [][]byte

	for _, n := range data {
		node := CalHashSha256(n)
		mtns = append(mtns, node)
	}

	for {
		mtns = CalMerkleUpperHashs(mtns)
		if len(mtns) == 1 {
			break
		}
	}

	// bclogger.INFOMSGF("%x\n", mtns)

	return mtns[0]
}
