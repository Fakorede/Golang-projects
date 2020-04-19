package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	ReadSomething()
}

func ReadSomething() error {
	var r io.Reader = BadReader{errors.New("my awesome reader")}
	if _, err := r.Read([]byte("test something")); err != nil {
		fmt.Println("An error occured!", err)
		return err
	}

	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return 1, br.err
}
