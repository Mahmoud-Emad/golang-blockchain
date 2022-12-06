package main

import (
	"fmt"
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash 		[]byte
	Data		[]byte
	PrevHash 	[]byte
}

func (b *Block) DriveHash() {
	info 	:= bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash 	:= sha256.Sum256(info)
	b.Hash	 = hash[:]
}

func CreateBlock(data string, pervHash []byte) *Block {
	block 	:= &Block{[]byte{}, []byte(data), pervHash}
	block.DriveHash()
	return block
}

func (chain *BlockChain) AddBlock (data string){
	pervBlock 	:= chain.blocks[len(chain.blocks)-1]
	new 		:= CreateBlock(data, pervBlock.Hash)
	chain.blocks = append(chain.blocks, new) 
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main(){
	chain := InitBlockChain()

	// Adding some blocks
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}