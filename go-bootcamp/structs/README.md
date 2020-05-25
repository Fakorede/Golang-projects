# STRUCT

**Structs** allows us to represent concepts by declaring a blueprint for a set of related types in a single type.

Structs are used to shape our own types using other types.
Example, we can create a `Person struct` to represent a Person, a `Movie struct` for a Movie recommendation engine, or a `Rental struct` to create a rental application etc.

We can think of Struct as a Class in OOP languages. By using Structs, we can group related data about a concept in a single Struct type.

Structs are used as *blueprints* through which we can create as many values as we want in runtime. 

Unlike Arrays, Slices and Maps, we can store different types of data in a Struct.
**Struct fields** are fixed at *compile-time*, so we cannot change the field names and types in runtime. It also means we cannot add new fields in runtime. **Struct values** fill them in runtime.

Example Struct:

```
type VideoGame struct {
    Title string
    Genre string
    Published bool
}

pacman := VideoGame {
    Title: "Pac-man",
    Genre: "Arcade Game",
    Published: true,
}
```

`Title`, `Genre`, `Published` are called **Struct Fields**. While `Pac-man`, `Arcade Game`, `true` are the **Struct Values**.


## Initializing Struct Values

```
type person struct {
    name, lastname string
    age            int
}

var dicaprio person
var picasso person

dicaprio.name = "Leornado"
dicaprio.lastname = "DiCaprio"
dicaprio.age = 47

picasso.name = "Pablo"
picasso.lastname = "Picasso"
picasso.age = 91

fmt.Printf("\nDiCaprio: %#v", dicaprio)
fmt.Printf("\nPicasso: %#v", picasso)

fmt.Printf("\n\n%s's age is %d", picasso.lastname, picasso.age)
```

Struct Values can also be assigned alternatively like so:

```
picasso := person{name: "Pablo", lastname: "Picasso", age: 91}
```

## Comparing Struct Values

Struct values are bare types(just like arrays).

Struct values are copied along with their fields.

Two structs are equal if all their fields are equal.

Their types should be identical.

We cannot compare struct values if they contain fields with uncomparable types like slice, map, or function value.

```
type song struct {
    title, artist string
}

func main() {
    song1 := song{title: "Intentions", artist: "Justin Bieber"}
	song2 := song{title: "Perfect", artist: "Ed Sheeran"}

	song1 = song2

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	if song1 == song2 {
		fmt.Println("Songs are equal.")
	} else {
		fmt.Println("Songs are not equal.")
	}
}
```

## Embedding Structs

```
func main() {
    type text struct {
        title string
        words int
    }

    type book struct {
        text text
        isbn string
    }

    moby := book {
        text: text{title: "moby dick", words: 206052},
        isbn: "102030",
    }

    fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)
}
```

A struct field without a name is called an **Anonymous Field**.

```
type book struct {
    text
    isbn string
}

 moby := book {
    text: text{title: "moby dick", words: 206052},
    isbn: "102030",
}

fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

```

When the field name conflicts, the parent type takes priority.

But grouping variables in a structs, we have a more organized and less cluttered code.

### Encoding Values to JSON(Marshalling)

Why JSON?

Human Readable.

Computers can easily understand it.

Widely used and supported.

```
type user struct {
    username string
    password string
    permissions map[string]bool
}

func main() {
    users := []user{
        {"inanc", "1234", nil},
    }
}
```

## Properties

- Used to represent a "thing" with a lot of different properties

- Keys don't support indexing

- Need to know all the different fields at compile time

- Values can be of different type

- Value Type

