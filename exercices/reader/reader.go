package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buf []byte) (n int, err error) {
	// fill one byte at a time
	buf[0] = 'A'
	return 1, nil

	// fill the whole provided buffer
	/*
		for i := range buf {
			buf[i] = 'A'
		}
		return len(buf), nil
	*/
}

func main() {
	reader.Validate(MyReader{})
}
