package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var Blockchain []Block


type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    CurrentHash  string
}


func CalculateHash(stringToHash string) string {
    hash := sha256.New()
    hash.Write([]byte(stringToHash))
    return hex.EncodeToString(hash.Sum(nil))
}


func NewBlock(transaction string, nonce int, previousHash string) *Block {
    newBlock := Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
        CurrentHash:  CalculateHash(transaction + previousHash + fmt.Sprint(nonce)),
    }
    Blockchain = append(Blockchain, newBlock)
    return &newBlock
}


func ListBlocks() {
    for i, block := range Blockchain {
        fmt.Printf("Block %d:\n", i)
        fmt.Printf("Transaction: %s\n", block.Transaction)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Current Hash: %s\n\n", block.CurrentHash)
    }
}


func ChangeBlock(index int, newTransaction string) {
    if index < 0 || index >= len(Blockchain) {
        fmt.Println("Invalid block index.")
        return
    }

    
    Blockchain[index].Transaction = newTransaction

    
    Blockchain[index].CurrentHash = CalculateHash(Blockchain[index].Transaction + fmt.Sprint(Blockchain[index].Nonce) + Blockchain[index].PreviousHash)
}



func VerifyChain() {
	for i := 1; i < len(Blockchain); i++ {
		
		if Blockchain[i].PreviousHash != Blockchain[i-1].CurrentHash {
			fmt.Printf("Blockchain is invalid at block %d\n", i+1)
			return
		}
	}
	fmt.Println("Blockchain is valid.")
}
