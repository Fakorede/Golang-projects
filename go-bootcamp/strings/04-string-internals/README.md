# String Internals

## Why string values are immutable?

A string is a data structure that points to a read-only backing array.

**String Header** stores where a string value starts and ends in computer memory.

String Header
|         |     |                    |
| :-----: | --- | ------------------ |
| Pointer | 57  | -> where it starts |
| Length  | 5   | -> where it ends   |



Backing Array 
|       |     |     |     |     |
| :---: | --- | --- | --- | --- |
|   h   | e   | l   | l   | o   |

String values which are exactly the same share the same backing array for memory efficiency.

Slicing a string also reuses the same backing array in memory which makes it very efficient.

To make a change to a string, we'll need to convert to a []byte which allocates a new backing array and copies the string values.

Here is what the string header looks like in Go.

```
// StringHeader is used by a string value
// In practice, you should use: reflect.Header
type StringHeader struct {
	// points to a backing array's item
	pointer uintptr // where it starts
	length  int     // where it ends
}

// dump prints the string header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
```

