## Bytes, Runes, and Strings

- A string value is a series of bytes

```
"hey" => []byte{104, 101, 121}
```

- `string` and `[]byte` are interchangeably convertible.

```
[]byte("hey")
string([]byte{104, 101, 121})
```

- We can represent a character using a `rune literal`. Rune literals are typless numbers.

```
'h' => 104
'e => 101
'y' => 121
```

`Unicode Code Points` are `runes` ex 104, 101, 121 etc

- A `rune literal` can represent `byte(uint8)`, `rune(int32)`, or any other numeric type.

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

### Converting, Indexing and Slicing

- Strings are immutable(read-only) byte slices. We can only make a change by converting the string to a byte slice.

```
func main() {
	str := "Young"
	bytes := []byte(str)
	bytes[0] = 'N'
	bytes[1] = 'O'

	str = string(bytes)
	fmt.Printf("%s\n", str)
}
```

- `len` counts the bytes while `utf8.RuneCountInString` counts the runes.

    `RuneCountInString` and `RuneCount` are **mirror functions** that helps to prevent conversions btw []byte and string because it is costly(may allocate new memory). There are a lot of functions like these because []byte and string are interchangeable values.

```
func main() {
	str := "Yūgen ☯ 💀"

	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str)) // 15
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str)) // 9

	fmt.Printf("% x\n", bytes)
    fmt.Printf("\t%d bytes\n", len(bytes)) // 15
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes)) // 9
}
```

- **for range** loop jumps over the runes of a string, rather than its bytes. Each index returns the starting index of the next rune. UTF-8 is a variable length encoding(for efficiency). So, each rune may start at a different index.

```
    str := "Yūgen ☯ 💀"

	for i, r := range str {
	    fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}
```

- A string is a byte slice so we should be able to use it as a slice.

```
    str := "Yūgen ☯ 💀"
    fmt.Printf("1st byte    : %c\n", str[0])
    fmt.Printf("2nd byte    : %c\n", str[1]) // nt ok cos it has 2 bytes.
    fmt.Printf("2nd rune    : %s\n", str[1:3])
    fmt.Printf("last rune    : %s\n", str[len(str)-4:]) // has 4 bytes.
```

`Runes` in a UTF-8 encoded string can have a different number of bytes because utf-8 is a variable-length encoding.
Go doesn't allow us to index UTF-8 strings easily because of efficiency reasons.

- Theres an easier way, however, it is inefficient. `[]rune(string)` creates a new slice, and copies each rune to the new slice's backing array.

    Indexing and slicing a rune is very easy.

```
    str := "Yūgen ☯ 💀"
    runes := []rune(str)

    fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d runes\n", len(runes)) // 9
    fmt.Printf("1st rune   : %c\n", runes[0]) // Y
	fmt.Printf("2nd rune   : %c\n", runes[1]) // ū
	fmt.Printf("first five : %c\n", runes[:5]) // [Y ū g e n]

```

- Why Go doesn't allow us to use `[]rune` and forces us to use `[]byte` and `string`

```
    str := "Yūgen ☯ 💀"
    runes := []rune(str)

	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes)) // 36 bytes

```

The rune(int32) in the example has 9 elements and occupies 36bytes. However, when we use a string value, it only has 15 bytes. This is because the UTF-8 uses a variable-length scheme with each rune having different lengths.

Also, the Go standard lib and most libraries prefer to use `[]byte` instead of `[]rune`.



## Learn More

https://blog.golang.org/strings