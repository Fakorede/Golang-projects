# Public and Private Functions

Visibility is on a per package level only.

A package is the smallest unit of private encap­sulation in Go.

All identifiers defined within a package are visible throughout that package.
When importing a package you can access only its exported identifiers.

An identifier is exported if it begins with a capital letter. If an identifier starts with a lower case letter, it can only be accessed from within the package.

### Example

Defining the functions `Sum`, `Divide` and `Add` of the `expressions.go` starting with capital letters makes them accessible in `main.go`.

```
// Sum : sums up slice values provided as argument.
func Sum(values ...float64) float64 {

}
```

### Usage

Run the program with the command:

```
go run main.go expressions.go
```
