# Error Handling

_Error Handling_ is important because not everything goes as expected in our programs. We need to handle error so we can decide what to do when an error shows up.

A lot of things could go wrong like:

- Network problems.
- File access problems.
- User input problems etc.

**Nil value** means that the value is not initialized yet. It is extensively used for _Error Handling_ and means there is no error.

It is also a zero value for pointer-based types like Pointers, Slices, Maps, Interfaces and Channels.

### Error Handling Example

```

func main() {
    age := os.Args[1]

	n, err := strconv.Atoi(age)

	if err != nil { // checks whether there is an error or not
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}

```

When we call a function that returns an error, we should always check for the error immediately after the call to that function. This is to prevent unnecessary bugs when running our programs.

### Usage

```
go run main.go 50
```
