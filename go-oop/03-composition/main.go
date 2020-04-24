package main

import "fmt"

// CreditAccount struct
type CreditAccount struct{}

// AvailableFunds method for CreditAccount
func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit account funds...")
	return 250
}

// CheckingAccount struct
type CheckingAccount struct{}

// AvailableFunds method for CheckingAccount
func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit account funds...")
	return 125
}

// HybridAccount struct
type HybridAccount struct {
	CreditAccount
	CheckingAccount
}

// AvailableFunds method for HybridAccount
func (h *HybridAccount) AvailableFunds() float32 {
	return h.CreditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}

func main() {
	ha := &HybridAccount{}
	fmt.Println(ha.AvailableFunds())
}
