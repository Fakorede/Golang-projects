# Functions and Pointers

Functions are re-usable code blocks.

Pointers stores memory address.


Every function starts with a `func` keyword and has a body wrapped between curly braces.

- After the func keyword comes the name and the function may declare zero or more input parameters.

```
func name(inputParams) {}

```


- Each parameter should have a name and a type. Example:

```
func sum(a int, b int) {}

```

This can also be written as:

```
func sum(a, b int) {}

```

- If a function doesn't define any return values, then we can call the `return` statement without any arguments.

```
func Print() {
    return
}
```
- A function may declare zero or more result values

```
func name(inputParams) (resultParams) {

}
```

An example is

```
func Atoi(s string) (int, error) {
    return 0, nil
}
```

As a convention, an error result value should be always declared as the last parameter.

- We don't need to use parenthesis when a function declares a single result value.

An example is

```
func Itoa(i int) string {
    return "some string value"
}
```

- Input params and result values are copied.

- Package level variables can cause hard to debug problems later on. So, its best to confine variables with functions.

### Naked Returns

Returns _named result values_ automatically.

```
func parse(p parser, line string) (parsed result, err error) {
    //...

    return
}
```

Inside the function, the result values become variables.
