# lucky-number

This is a number guessing game.

### How the Game Works?

The program automatically generates 5 numbers between 0 and the supplied number. You have a winning guess if the number you supplied is one of the generated numbers.

### Code specific

```
time.Now() - current time as a Time value
t.UnixNano() - unix time in nanoseconds
rand.Seed() - seeds the random number algorithm
rand.Intn() - returns a random number depending on the seed value
```

### Usage

```
go run main.go 5
```

> where 5 is your guess/supplied number.
