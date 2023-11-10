package main

import (
	"log"

	"github.com/gmira/blockchain/src/api/domain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := domain.NewBlockchain()

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(10, previousHash)

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)

	blockchain.Print()
}
