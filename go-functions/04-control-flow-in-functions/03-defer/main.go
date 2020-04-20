package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	if err := ReadFullFile(); err != nil {
		fmt.Printf("something bad occured: %s", err)
	}
}

// ReadFullFile : ReadFullFile
func ReadFullFile() error {
	var r io.ReadCloser = &SimpleReader{}

	defer func() {
		_ = r.Close()
	}()

	defer func() {
		fmt.Println("before for-loop")
	}()

	for {
		value, err := r.Read([]byte("this text does nothing"))
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

// ReadSomethingBad ...
func ReadSomethingBad() error {
	var r io.Reader = &BadReader{errors.New("my bad reader")}
	value, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}

	fmt.Println(value)

	return nil
}

// BadReader : BadReader
type BadReader struct {
	err error
}

func (br *BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

// SimpleReader : SimpleReader
type SimpleReader struct {
	count int
}

// Read : Read
func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count > 3 {
		return 0, io.EOF //errors.New("random error") //io.EOF
	}
	br.count++

	return br.count, nil
}

// Close : Close
func (br *SimpleReader) Close() error {
	fmt.Println("Closing reader...")
	return nil
}
