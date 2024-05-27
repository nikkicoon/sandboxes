package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}


// Type to represent the Blockchain
type BlockChain struct {
	// Array to pointers of Block
	Blocks []*Block
}

//func (b *Block) DeriveHash() {
//	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//	hash := sha256.Sum256(info)
//	b.Hash = hash[:]
//}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run() // execute the Run function on the pow
	return block
}

// XXX: what is the leading (...) in this function?
func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks) - 1]
	// create current block
	new := CreateBlock(data, prevBlock.Hash)
	// append this block to our chain
	chain.Blocks = append(chain.Blocks, new)
}

// create first block ("Genesis Block")
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
