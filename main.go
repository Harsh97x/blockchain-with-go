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
}
