package paybroker

import "fmt"

// PaymentBrokerAccount struct
type PaymentBrokerAccount struct{}

// ProcessPayment for PaymentBrokerAccount
func (p *PaymentBrokerAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment with a payment broker...")
	return true
}
