package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(array []byte) (int, error) {

	slice := array[:cap(array)]

	for index := range slice {
		slice[index] = 'A'
	}

	return cap(slice), nil
}

func main() {
	reader.Validate(MyReader{})
}
