package main

import (
	"blockchain/block"
	"blockchain/wallet"
	"fmt"
)

func main() {
	w := wallet.NewWallet()
	w1 := wallet.NewWallet()

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.Address(), w1.Address(), 1.0)

	mergenChain := block.NewBlockchain(w.Address())
	isAdded := mergenChain.AddTransaction(
		w.Address(),
		w1.Address(),
		1.0,
		w.PublicKey(),
		t.GenerateSignature())
	fmt.Println("Added?", isAdded)

	mergenChain.Mining()
	mergenChain.Print()

	fmt.Printf("A %.1f\n", mergenChain.CalculateTotalAmount(w.Address()))
	fmt.Printf("B %.1f\n", mergenChain.CalculateTotalAmount(w1.Address()))
}
