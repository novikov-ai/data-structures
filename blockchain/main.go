package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const initHash = "00000"

func main() {
	blockchain := New("test", 1)
	println(blockchain.blocks)

	blockchain.Add("my first block")
	blockchain.Add("my second block")

	println(HashWithNonce(*blockchain.last, 5))
	print()
}

type Block struct {
	index int
	nonce int
	data  string
	hash1 string // current block
	hash2 string // prev block
	time  time.Time
}

type BlockChain struct {
	blocks     []Block
	index      map[string]int
	last       *Block
	nonceCount int
}

func New(data string, nonceCount int) BlockChain {
	block := Block{
		index: 0,
		hash1: initHash,
		hash2: "",
		data:  data,
		time:  time.Now(),
	}

	return BlockChain{
		blocks: []Block{block},
		index: map[string]int{
			block.hash1: 0,
		},
		last:       &block,
		nonceCount: nonceCount,
	}
}

func (bc *BlockChain) Add(data string) {
	nb := Block{
		index: bc.last.index + 1,
		nonce: 0, // ???
		hash2: bc.last.hash1,
		data:  data,
		time:  time.Now(),
	}

	nb.hash1 = Hash(nb)

	bc.blocks = append(bc.blocks, nb)
	bc.index[nb.hash1] = nb.index
	bc.last = &nb
}

func Hash(b Block) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v%s%v%v", (b.index), b.data, b.nonce, b.time))) // hash based on 3 field
	hashValue := hasher.Sum([]byte(b.hash2))                                          // append hash of prev block
	return hex.EncodeToString(hashValue)
}

func HashWithNonce(b Block, nonceCount int) string {
	for {
		hashWithNonce := Hash(b)
		ok := checkHashForNonceCount(hashWithNonce, nonceCount)
		if ok {
			return hashWithNonce
		}

		b.nonce++
	}
}

func checkHashForNonceCount(hash string, nonce int) bool {
	if hash == "" {
		return false
	}

	hits := 0
	for i := len(hash) - 1; i >= 0; i-- {
		if hits >= nonce {
			return true
		}

		if hash[i] == '0' {
			hits++
			continue
		} else {
			return false
		}
	}

	return hits >= nonce
}

func (b *BlockChain) FindByHash(h string) (Block, bool) {
	index, ok := b.index[h]
	if !ok {
		return Block{}, false
	}

	return b.blocks[index], true
}

func (b *BlockChain) ValidateChain() bool {
	// метод, проверяющий целостность цепочки (корректность каждого хеша).
	prevHash := initHash
	for i, block := range b.blocks {
		if i == 0 {
			if block.hash1 != initHash {
				return false
			}
			if block.hash2 != "" {
				return false
			}
			continue
		}

		if block.hash1 != Hash(block) {
			return false
		}
		if block.hash2 != prevHash {
			return false
		}
		prevHash = block.hash1
	}

	return true
}