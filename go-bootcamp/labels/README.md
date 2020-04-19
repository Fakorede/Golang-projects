# Labels in GO

Labeled Statements in Go allows us to jump to a label using a `break`, `continue` or `goto` statement.

Labels do not conflict with names of other variables.

```
	var queries string
	_ = queries

queries:
	for _, q := range query {
		for i, w := range words {
			...
		}
	}

```

## Labeled Break

The `labeled break` allows us break from any `labeled loop`.

The labeled break works with a Switch statement as well.

### Example

```
queries:
	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break queries
			}
		}
	}
}
```

### Output

```
#4 : "again"
```

## Labeled Continue

The `labeled continue` continues from the `labeled loop` (the parent loop in this case) instead of the nested loop.

### Example

```
queries:
	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				continue queries
			}
		}
	}
}
```

### Output

```
#4 : "again"
#5 : "and"
```

## Break from a Switch with Labeled Break

We want to filter out unwanted words from the provided query words so it won't be searched for or appear in the search results.

### Example

In the example below, whenever the queried words matches the words in the case clause, the case clause is executed.

```
queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch q {
				case "and", "or", "the":
				// break
				break search
			}

			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				continue queries
			}
		}
	}
}
```

### Output

Using a normal break only breaks from the switch and continues executing the if statement.

```
#4 : "again"
#5 : "and"
```

To fix this, we add a `labeled break` statement to not only break from the switch but also from the nested loop.

```
#4 : "again"
```

### Usage

```
go run main.go again and
```

> where again & and are the arguments/query
