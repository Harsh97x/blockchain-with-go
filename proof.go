package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/magefile/mage/target"
)

const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	pow := &ProofOfWork{b, target}
	return pow
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) InitNonce (nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.Prevhash,
		pow.Block.Data,
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty)),
	},
		[]byte{},
	)
	return data
}

func ()
