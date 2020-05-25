# JSON ENCODING

JSON is a structured data exchange format. We can encode structs, maps and slices to JSON. In Go, this is called **Marshalling**. Decoding is called **Unmarshalling**.


We use JSON because:

- It is Human readable.
- Computers can easily understand it.
- Widely used and supported.


## Encoding values to JSON

### Example

Encode list of users to JSON.


```
type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}
```

***Note:** JSON package only encodes exported fields

`json:"username"` encodes the field as username.

`json:"-"` encodes the field but excludes it.

`json:"perms,omitempty"` encodes the field as perms and prevents encoding if it has a zero-value.

`json:",omitempty"` only prevents encoding if it has a zero-value.

```
func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"fab", "42", permissions{"admin": true}},
		{"dave", "644", permissions{"write": true}},
	}

	// json.Marshal(users)
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
```

`json.Marshal` encodes a value to JSON. `json.MarshalIndent` is an alternative to `json.Marshal` which formats the output to indent it.

```
go run main.go > users.json
```

This command saves the output to a file named `users.json`.

### Field Tags

Field Tags associates a static string metadata to a field. They are part of a struct's type and are compile-time constants. 

They are mostly used for controlling the encoding/decoding behavior.

We enclose Field Tags in a raw string literal because they usually include double-quotes.

We denote a Field tag as a key-value. The key denotes a package name which means only the package(in this case JSON package) will read the tag. The value may contain multiple options and are conventionally enclosed in double-quotes seperated by a comma.

We use the options to tell the json package how to encode and decode a field.

## Decoding values from JSON

### Example

In the program below, we read the json data into a `[]byte` and then decode it into a `[]user`

```
type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"permissions"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}
```

`json:"username"` tells the json package to get the data from the username field into the Name field.

`json.Unmarshal` decodes from JSON to a value.

`&` finds the memory address of the `users` variable which it passess to the Unmarshal function.