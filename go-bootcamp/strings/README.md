## Bytes, Runes, and Strings

A string value is a series of bytes

```
"hey" => []byte{104, 101, 121}
```

`string` and `[]byte` are interchangeably convertible.

```
[]byte("hey")
string([]byte{104, 101, 121})
```

We can represent a character using a `rune literal`.

```
'h' => 104
'e => 101
'y' => 121
```

`Unicode Code Points` are `runes` ex 104, 101, 121 etc

A `rune literal` can represent `byte(uint8)`, `rune(int32)`, or any other `integer` type

```
0b01101000        <-- 'h' -->            0x68
(binary)           104 (decimal)        (hexadecimal)
```

```
104 = 64 + 32 + 8 = 2^6 + 2^5 + 2^3 (base 2)
104 = 1 * 10^2 + 0 * 10^1 + 4 * 10^0 (base 10)
104 = 0x68 = 6 * 16^1 + 8 * 16^0 = 96 + 8 (base 16)
```

A `rune literal` or `character` or `code point` is a an integer number.

`Rune literals` are typeless. They can be assigned to a variable without any numberic type.


### Character-set Program

```
func main() {
    start, stop := 'A', 'Z'

    for n := start; n <= stop; n++ {
        fmt.Printf("%c => %[1]d\n", n)
    }
}
```


## Learn More

https://blog.golang.org/strings