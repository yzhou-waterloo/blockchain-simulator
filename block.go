package main

import (
	"time"
)

// type Block keeps block headers
type Block struct {
	Timestamp     int64  // When the block was created
	Data          []byte // Transactions
	PrevBlockHash []byte // Hash of the previous block
	Hash          []byte // Hash of this block’s contents
	Nonce         int    // Nonce we used in the mining process of the current block
}

// NewBlock(data, prevBlockHash) creates a new block that contains data and prevBlockHash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
