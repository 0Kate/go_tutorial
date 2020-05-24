package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (r MyReader) Read([]byte) (int, error) {
    return int('A')
}

func main() {
    reader.Validate(MyReader{})
}
