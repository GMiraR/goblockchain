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
	blockchain.CreateBlock(5, "hash test 1")
	blockchain.CreateBlock(2, "hash test 2")
	blockchain.CreateBlock(10, "hash test 3")
	blockchain.Print()
}
