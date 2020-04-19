# Anonymous Functions

An **Anonymous function** is a function that is declared without any named identifier to refer to it.

It is useful when we want to create an inline function. In Go, an **Anonymous function** can form a closure. An anonymous function is also known as _function literal_.

```
func main() {
    func(){
        fmt.Println("My first anonymous function...")
    }()
}
```

In Go, you are allowed to assign an **Anonymous function** to a variable. When you assign a function to a variable, then the type of the variable is of function type and you can call that variable like a function call:

```
func main() {

    // Assigning an anonymous function to a variable
    value := func(){
        fmt.Println("Welcome to reality!")
    }

    value()

}
```

**Anonymous functions** can accept inputs and return outputs, just as standard functions do.

```
func main() {
	v := func(name string) string {
		return name
	}

	val := v("John Doe")
	fmt.Println("My name is", val)

}
```

```
func main() {
	func(name string){
		fmt.Println("My name is", name)
	}("Fab")
}
```
