package main

import (
	"github.com/fakorede/learning-golang/go-oop/04-polymorphism/paybroker"
	"github.com/fakorede/learning-golang/go-oop/04-polymorphism/payment"
)

// Option interface
type Option interface {
	ProcessPayment(float32) bool
}

func main() {
	var option Option

	option = &payment.CreditCard{}
	option.ProcessPayment(500)

	option = &payment.CheckingAccount{}
	option.ProcessPayment(500)

	option = &paybroker.PaymentBrokerAccount{}
	option.ProcessPayment(500)
}
