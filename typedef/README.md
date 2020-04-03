# Type Definition

Type Definition allows us create our own types in Go.

> This program aims to demonstrate what defined types are and how to create them. It also explains underlying types and aliased.

### Explanation

##### Defined Types

Go allows us to create our own types using Type Definition. These created types which aren't Predefined(e.g. int, float, bool) are known as _Defined Types_.

An example is the _Duration_ type defined in the [time](https://golang.org/pkg/time/#Duration). package.

It is defined using the syntax `type Duration int64`.

Here, `Duration` is the name of the new type based on int64. `int64` represents the _Underlying Type_.

- Defined Types ensures Type Safety i.e. they cannot be used with other types without doing a type conversion.
- Methods can be attached to Defined Types as in the [time.Duration](https://golang.org/pkg/time/#Duration) package.
- Values of a defined type can be converted to another type if that type's underlying type is identical.

##### Underlying Types

_Underlying Types_ are used to define the structure of a new type. Take the example below:

```
type Duration int64
type MyDuration Duration
type MoreDuration MyDuration
```

In the declarations above, `int64` is the _Underlying type_ for the _Defined Types_ `Duration`, `MyDuration`, and `MoreDuration`.

These types have a representation of signed integers and a byte size of 8 which they all get from `int64`.

- A defined type shares the same underlying type with its source type. Ex. `MoreDuration` shares the same underlying type `int64` as `MyDuration`.
- Types with identical underlying types can be converted to each other. Ex. `MyDuration` can be converted to `Duration` and `Duration` can be converted to `MoreDuration`.

**Note:** We can deduce from the above that there are no type hierarchies in Go.

##### Aliased Types

These are the very same types with a new name.

**Example**

- `byte` and `uint8` are exactly the same types with different names.
- `rune` and `int32` are also exactly the same types with different names.

- _Aliased Types_ can be used together without type conversions.
- _Aliased Types_ can be represented as:

```
type byte = uint8
type rune = int32
```

Rune is used for representing unicode characters.

### Run Program

```
$ go run main.go definedtypes.go underlyingtypes.go
```
