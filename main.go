package main

import (
	"assignment01bca/assignment01bca"
	"fmt"
)

func main() {
	
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	fmt.Printf("Genesis Block Hash: %s\n", genesisBlock.CurrentHash) 

	
	assignment01bca.NewBlock("Alice sends 1 BTC to Bob", 12345, genesisBlock.CurrentHash)
	assignment01bca.NewBlock("Bob sends 0.5 BTC to Charlie", 54321, assignment01bca.Blockchain[1].CurrentHash)

	
	assignment01bca.ListBlocks()

	
	assignment01bca.ChangeBlock(1, "Alice sends 5 BTC to Eve")

	
	assignment01bca.ListBlocks()

	
	assignment01bca.VerifyChain()
}
