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

	blockchain.AddTransaction("Bilbo", "Frodo", 20.0)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(10, previousHash)

	previousHash = blockchain.LastBlock().Hash()
	blockchain.AddTransaction("Aragorn", "Gollum", 25.5)
	blockchain.AddTransaction("Gandalf", "Legolas", 10.0)
	blockchain.CreateBlock(2, previousHash)

	blockchain.Print()
}
