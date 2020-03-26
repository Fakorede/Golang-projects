# Hello World

> A basic hello world program in the GO programming lang.

### Explanation

**package main** - Each source file begins with a package declaration. Here _package main_ states which package the file belongs to.

**import "fmt"** - This line imports the **fmt** package which contains functions for printing formatted output and scanning input. **Println** is one of those basic functions. It prints one or more values, seperated by spaces, with a newline character at the end so that the values appear as a single line of input.

**func main** - This is where execution of the program begins. Whatever **main** does is what the program does.
**func** is a function declaration. A function declaration consists of the keyword **func**, the name of the fucntion, a parameter list(empty for main), the body of the function and the statement that defines what it does(enclosed in braces).

### Run Program

```
$ go run helloworld.go
```
