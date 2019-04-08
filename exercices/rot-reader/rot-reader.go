package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = r13.r.Read(buf)
	if err != nil {
		return
	}
	for i, c := range buf {
		buf[i] = rot13(c)
	}
	return n, nil
}

func rot13(c byte) byte {
	if c >= 'a' && c <= 'm' ||
		c >= 'A' && c <= 'M' {
		return c + 13
	}
	if c >= 'n' && c <= 'z' ||
		c >= 'N' && c <= 'Z' {
		return c - 13
	}
	return c

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
