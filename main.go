package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	type Block struct {
		Hash     []byte
		Data     []byte
		Prevhash []byte
	}

	type Blockchain struct {
		blocks []*Block
	}

	func (b *Block) derivehash () {
		info := bytes.Join([][]byte{b.Data, b.Prevhash}, []byte{})
		hash := sha256.Sum256(info)
		b.hash = hash[:]
	}

	func CreateBlock(data string, prevHash []byte) *Block {
		block := &Block{[]byte{}, []byte(data), prevHash}
		block.DeriveHash()
		return block
	}

	func (chain *Blockchain) AddBlock (data string) {
		prevblock := chain.blocks[len(chain.block-1)]
		new := Createblock(data, prevblock)
		chain.blocks := append(chain.blocks, new)
	}

	func Genesis() *Block {
		return CreateBlock("Genesis", []byte{})
	} 
}
