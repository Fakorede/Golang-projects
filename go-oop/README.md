# Object-Oriented Programming with Go

## Encapsulation - not knowing the implementation details

Encapsulation is about accessing a service on an object without knowing how that service is implemented.

### Challenges of Encapsulation in Go

- No classes
- No concept of private data

### Strategies

- Package Oriented Design

```
package payment

type CreditAccount struct {
    accountNumber string
    accountOwner string
}

func (c CreditAccount) AccountNumber() string
func (c CreditAccount) AccountOwner() string
func (c CreditAccount) AvailableCredit() float32

```

- Interfaces

Anything that'll be interacting with `CreditAccount`, will be doing so through this interface.

```
type Payment interface {
    AccountNumber() string
    AvailableCredit() float32
}

type Option interface {
	ProcessPayment(float32) bool
}

```

## Massage Passing - how to handle request

Involves sending a message to an object, but letting that object determine what to do with it.

Its not just about defining a method on an object and invoking the object. Its about abstracting away the call from the sender and what services they call.

The caller don't know which service is invoked. It just knows there's a method available on the interface it can call.

### Strategies

- Interfaces

```

type Option interface {
    ProcessPayment(float32) bool
}

type CashAccount struct{}
func(c *CashAccount) ProcessPayment(amount float32) bool {...}

type CreditAccount struct{}
func(c *CreditAccount) ProcessPayment(amount float32) bool {...}

// ----------Client------------------
var paymentOption Option
paymentOption = &CashAccount{}                      - message passing
ok := paymentOption.ProcessPayment(500)

paymentOption = &CreditAccount{}                    - message passing
ok := paymentOption.ProcessPayment(500)
```

- Channels

```
type CreditAccount struct{}

// package scoped method/service
func (c *CreditAccount) processPayment(amount float32) {
    fmt.Println("Processing credit card payment")
}

// constructor that takes in a Channel
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
    creditAccount := &CreditAccount{}
    // go routine listens for msg on the Channel
    go func(chargeCh chan float32) {
        for amount :=range chargeCh {
            // and invokes the service it calls
            creditAccount.processPayment(amount)
        }
    }(chargeCh)

    return creditAccount
}

// ----------Client------------------
cahrgeCh := make(chan float32)
account := CreateCreditAccount(chargeCh)
chargeCh <- 500
```

## Composition - how to reuse exisiting behavoirs


Behavior reuse strategy where a type contains objects that have desired functionality. The type delegates calls to those objects to use their behaviors.

### Type Embedding

```
type Account struct{}

func (a *Account) AvailableFunds() float32 {}
func (a *Account) ProcessPayment(amount float32) bool {}


type CreditAccount struct{
    Account
}


ca := &CreditAccount{}

funds := ca.AvailableFunds()

```

## Polymorphism

The abilitiy to substitue a family of types that implement a common set of behaviors.

### Strategy - Interfaces

```
type Option interface {
	ProcessPayment(float32) bool
}

type CreditAccount struct {...}
func (c *CreditAccount) ProcessPayment(amount float32) bool {...}
```