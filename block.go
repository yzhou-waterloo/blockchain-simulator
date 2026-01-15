package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// type Block keeps block headers
type Block struct {
	Timestamp     int64  // When the block was created
	Data          []byte // Transactions
	PrevBlockHash []byte // Hash of the previous block
	Hash          []byte // Hash of this block’s contents
}

// SetHash() combines header information to calculate the current hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock(data, prevBlockHash) creates a new block that contains data and prevBlockHash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
