// TODO parei no episódio 2
// https://www.youtube.com/watch?v=aE4eDTUAE70&list=PLpP5MQvVi4PGmNYGEsShrlvuE2B33xV1L&index=2

package main

import (
	"fmt"
	"furycoin/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("\nPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)

		isValid := pow.Validate()
		fmt.Printf("Pow: %s\n", strconv.FormatBool(isValid))
	}
}
