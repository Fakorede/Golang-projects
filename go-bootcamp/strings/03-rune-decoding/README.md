How to get runes from an utf-8 encoded string

```
func main() {
    text := `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
    Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
    Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

    for i := 0; i < len(text); {
        r, size := utf8.DecodeRuneInString(text[i:])
        fmt.Printf("%c", r)

        i += size
    }

    fmt.Println()

}
```

for range loops automatically decode the runes in a string value.
```
func main() {
    text := `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
    Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
    Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

    for _, r := range text {
        fmt.Printf("%c", r)
    }

    fmt.Println()
}
```

Make the first rune uppercase
```
func main() {
    word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

    // 1st way: `for range`
	// you can't get the runes by range overing a byte slice
	// first, you need to convert it to a string
	var size int
	for i := range string(word) {
		if i > 0 {
			size = i
			break
		}
	}

    // 2nd way: let's do it using the utf8 package's DecodeRune function
	_, size = utf8.DecodeRune(word)

    // overwrite the current bytes with the new uppercased bytes
	copy(word[:size], bytes.ToUpper(word[:size]))

	// to get printed bytes/runes need to be encoded in a string
	fmt.Printf("%s = % [1]x\n", word)

}

```