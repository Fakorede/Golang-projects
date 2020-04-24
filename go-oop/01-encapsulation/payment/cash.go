package payment

import "fmt"

// Cash struct
type Cash struct {
}

// CreateCashAccount CreateCashAccount *Cash
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment ProcessPayment(amount float32) bool
func (c *Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
