package main

import "fmt"

// CreditAccount struct
type CreditAccount struct{}

// package scoped method/service
func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment")
}

// CreateCreditAccount constructor that takes in a Channel
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	// go routine listens for msg on the Channel
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			// and invokes the service it calls
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

func main() {
	chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)
}
