# Input Scanning

### Bufio.Scanner

Scans an input stream line by line into a buffer

Syntax:

```
in := bufio.NewScanner(os.Stdin)
```

#### Methods

in.Scan()
in.Text()
in.Bytes()
in.Err()

Run the program:

```
go run main.go < proverbs.txt
```

# Use Maps as Sets

A set contains only unique values.