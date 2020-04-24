package main

import (
	"github.com/fakorede/learning-golang/go-oop/gopay/payment"
)

func main() {
	// credit := payment.CreateCreditAccount(
	// 	"Debra Ingram",
	// 	"1111-2222-3333-4444",
	// 	5,
	// 	2022,
	// 	123,
	// )

	// fmt.Printf("Owner name: %v\n", credit.OwnerName())
	// fmt.Printf("Card number: %v\n", credit.CardNumber())
	// fmt.Println("Trying to change the card number")
	// err := credit.SetCardNumber("invalid")
	// if err != nil {
	// 	fmt.Printf("That didn't work: %v\n", err)
	// }
	// fmt.Printf("Available credit: %v\n", credit.AvailableCredit())

	var option payment.Option

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2022,
		123,
	)

	option.ProcessPayment(500)

	option = payment.CreateCashAccount()

	option.ProcessPayment(500)
}
