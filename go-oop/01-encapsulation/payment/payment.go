package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Option PaymentOption interface
type Option interface {
	ProcessPayment(float32) bool
}

// CreditCard struct
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditAccount func
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// ProcessPayment ProcessPayment(amount float32) bool
func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

// OwnerName OwnerName() string
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName SetOwnerName(value string) error
func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}

	c.ownerName = value
	return nil
}

// CardNumber CardNumber() string
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber SetCardNumber(value string) error
func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card format")
	}

	c.cardNumber = value
	return nil
}

// ExpirationDate ExpirationDate() (int, int)
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate SetExpirationDate(month, year int) error
func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}

	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode SecurityCode() int
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode SetSecurityCode(value int)
func (c *CreditCard) SetSecurityCode(value int) {

}

// AvailableCredit AvailableCredit() float32
func (c CreditCard) AvailableCredit() float32 {
	return 5000 // can come frm a webservice, clien't doesn't know or care
}
