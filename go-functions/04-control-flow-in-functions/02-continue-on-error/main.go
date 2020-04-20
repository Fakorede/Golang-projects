package main

import (
	"fmt"
	"io"
)

func main() {
	if err := ReadFullFile(); err == io.EOF {
		fmt.Println("read file successfully")
	} else if err != nil {
		fmt.Println("something went wrong")
	}
}

func ReadFullFile() error {
	var r io.Reader = &SimpleReader{}
	for {
		value, err := r.Read([]byte("test something"))
		if err == io.EOF {
			fmt.Println("finished reading file, breaking out of loop...")
			break
		} else if err != nil {
			return err
		}

		fmt.Println(value)
	}

	return nil
}

type SimpleReader struct {
	count int
}

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count++

	return br.count, nil
}
