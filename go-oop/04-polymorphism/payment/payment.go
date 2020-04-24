package payment

import "fmt"

// // Option interface
// type Option interface {
// 	ProcessPayment(float32) bool
// }

// CreditCard struct
type CreditCard struct{}

// ProcessPayment for CreditCard
func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with a credit card...")
	return true
}

// CheckingAccount struct
type CheckingAccount struct{}

// ProcessPayment for CheckingAccount
func (c *CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with a check...")
	return true
}
