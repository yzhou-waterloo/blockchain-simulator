package main

type Blockchain struct {
	blocks []*Block
	// we will add a map later to find block by its hash
}

// AddBlock(data) adds a new block that contains the given transaction info
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock() creates the first block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewGenesisBlock() creates a new Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
