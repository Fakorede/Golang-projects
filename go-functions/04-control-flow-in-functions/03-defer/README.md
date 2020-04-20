# defer

In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.

Or in other words, defer function or method call arguments evaluate instantly, but they execute until the nearby functions returns.

You can create a deferred method, or function, or anonymous function by using the **defer** keyword.

```
// Function
defer func func_name(parameter_list Type)return_type{
// Code
}

// Method
defer func (receiver Type) method_name(parameter_list){
// Code
}

defer func (parameter_list)(return_type){
// code
}()
```

- In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order as shown in Example 2.

- In the defer statements, the arguments are evaluated when the defer statement executed, not when they called.

- Defer statements are generally used to ensure that the files are closed when your work is finished with them, or to close the channel, or to catch the panics in the program.

### Example 1

```
package main

import "fmt"

// Functions
func mul(a1, a2 int) int {

    res := a1 * a2
    fmt.Println("Result: ", res)
    return 0
}

func show() {
    fmt.Println("Hello!")
}

// Main function
func main() {

    // Calling mul() function here. mul function behaves like a normal function
    mul(23, 45)

    // Calling mul()function using defer keyword here the mul() function is defer function
    defer mul(23, 56)

    // Calling show() function
    show()
}
```

Output:

```
Result:  1035
Hello!
Result:  1288
```

```
### Example 2

// Go program to illustrate multiple defer statements
package main

import "fmt"

// Functions
func add(a1, a2 int) int {
    res := a1 + a2
    fmt.Println("Result: ", res)
    return 0
}

// Main function
func main() {

    fmt.Println("Start")

    // Multiple defer statements executes in LIFO order
    defer fmt.Println("End")
    defer add(34, 56)
    defer add(10, 10)
}
```

Output:

```
Start
Result:  20
Result:  90
End
```
