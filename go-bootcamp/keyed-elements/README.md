# Keyed Elements

_Keyed Elements_ defines the indexes of an array manually.

Each key corresponds to an index of the array

### Examples

```
rates := [...]float64{
    0: 2.5,
    1: 3.5,
    2: 4.5,
}

fmt.Printf("%#v", rates)
```

Outputs [3]float64{2.5, 3.5, 4.5}

```
rates := [...]float64{
    5: 1.5,
}

fmt.Printf("%#v", rates)
```

Outputs [6]float64{0, 0, 0, 0, 0, 1.5}

```
rates := [...]float64{
    5: 1.5,
    2.5,
    0: 0.5,
}

fmt.Printf("%#v", rates)
```

Outputs [7]float64{0.5, 0, 0, 0, 0, 1.5, 2.5}

### Usage

```
$ go run main.go
```

### Link to GO Playground

https://play.golang.org/p/KJxqQR0WWO5
