# Goto Statement

Allows us to jump to a label within the same function. We can also use the break or continue to jump to a label. However, they only work when used inside a for or switch statement.

On the otherhand, `goto` statement doesn't have that restriction. It can be used from almost anywhere as long as we're inside the same function.

### Example

The `goto` statement as well as the loop label makes the if statement behave like a real loop.

```
loop:
	if i < 3 {
        fmt.Println("looping...")
		i++
		goto loop
	}

    fmt.Println("done!")
}
```

### Output

```
looping...
looping...
looping...
done!
```

### Usage

```
go run main.go
```
