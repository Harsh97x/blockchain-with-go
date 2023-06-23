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
	buff := new(byte.Buffer)
	err := binary.Write()
}
