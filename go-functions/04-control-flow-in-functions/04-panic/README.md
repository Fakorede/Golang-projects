# Panic

In Go, **panic** is just like an exception, it also arises at runtime. Or in other words, panic **means** an unexpected condition arises in your Go program due to which the execution of your program is terminated.

Sometimes panic occurs at runtime when some specific situation arises like out-of-bounds array accesses, etc. or sometimes it is deliberately thrown by the programmer to handle the worst-case scenario in the Go program with the help of _panic()_ function.

The **panic** function is an inbuilt function which is defined under the built-in package of the Go language. This function terminates the flow of control and starts panicking.

**Syntax:**

```
func panic(v interface{})
```

It can receive any type of argument.

When the panic occurs in the Go program the program terminates at runtime and in the output screen an error message as well as the stack trace till the point where the panic occurred is shown.

Generally, in Go language when the panic occurs in the program, the program does not terminate immediately, it terminates when the go completes all the pending work of that program.

**Note**: Defer statement or function always run even if the program panics. See [example](#example) below.

**Usage of Panic:**

- You can use panic for an unrecoverable error where the program is not able to continue its execution.
- You can also use panic if you want an error for a specific condition in your program.

### Example

```
// Go program which illustrates the concept of Defer while panicking
package main

import (
    "fmt"
)

func entry(lang *string, aname *string) {

    // Defer statement
    defer fmt.Println("Defer statement in the entry function")

    // When the value of lang is nil it will panic
    if lang == nil {
        panic("Error: Language cannot be nil")
    }

    // When the value of aname is nil it will panic
    if aname == nil {
        panic("Error: Author name cannot be nil")
    }

    // When the values of the lang and aname are non-nil values it will print normal output
    fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname)
}

// Main function
func main() {

    alang := "GO Language"

    // Defer statement
    defer fmt.Println("Defer statement in the main function")

    // Here in entry function, we pass one non-nil and one-nil value
    // Due to nil value this method panics
    entry(&alang, nil)
}
```
