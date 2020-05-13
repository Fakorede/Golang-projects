package main

import (
	"fmt"
	"os"
)

/*
#1- Get and check the input
#2- Create a byte buffer and use it as the output
#3- Write input to the buffer as it is and print it
#4- Detect the link
#5- Mask the link
#6- Stop masking when whitespace is detected
#7- Put a http:// prefix in front of the masked link
*/

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide an argument")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)

		in bool
	)

	// don't use a for range loop because we want to loop over the bytes of the text, not its runes.
	for i := 0; i < size; i++ {

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			// fmt.Println(text[i : i+nlink])
			in = true
			buf = append(buf, link...)
			i += nlink
		}
		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}

		buf = append(buf, c)

	}

	fmt.Printf("%s", buf)

}
