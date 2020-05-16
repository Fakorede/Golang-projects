# Input Scanning

## Scanning User Input

### Bufio.Scanner

Scans an input stream line by line into a buffer

Syntax:

```
in := bufio.NewScanner(os.Stdin)
in.Scan()
```

### Methods

in.Scan()
in.Split()
in.Text()
in.Bytes()
in.Err()

### Example


```
func main() {

	os.Stdin.Close() // closes Standard Input
	in := bufio.NewScanner(os.Stdin)

	var lines int

	in.Scan()

	fmt.Println("Scanned Text :", in.Text())
	fmt.Println("Scanned Text :", in.Bytes())

	for in.Scan() {
		lines++
		fmt.Println("Scanned Text :", in.Text())
	}

	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
```

Run the program:

```
go run main.go < proverbs.txt
```



## Using Maps as Sets

A set contains only unique values.

```
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide a search word")
		return
	}

	query := args[0]

	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			// makes the map a set
			words[word] = true
		}
	}

	// for word := range words {
	// 	fmt.Println(word)
	// }

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}

	fmt.Printf("The input does not contain %q.\n", query)

}
```

Look for word in a file:
```
go run main.go pay < shakespaeare.txt
```

## Log Parser

Analyze and report website visits from a log file.