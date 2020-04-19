# Naming Return Values

Go enables us to name our return values. With this, we can assign our return values to variables and then simply `return`.

We can therefore re-write our code as:

```
func divide(p1, p2 float64) (ans float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}

	ans = p1 / p2

	return
}
```

### Usage

Run the program with the command:

```
go run main.go
```
