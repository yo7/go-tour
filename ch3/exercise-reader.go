package main

import (
	"fmt"
	"io"
)

type MyReader struct {
}

func (r MyReader) Read(bs []byte) (n int, err error) {
	bs = bs[:cap(bs)]
	for i := range bs {
		bs[i] = 'A'
	}
	return cap(bs), nil
}

func main() {
	bs := make([]byte, 8)
	r := MyReader{}
	for {
		n, err := r.Read(bs)
		fmt.Printf("%q\n", bs[:n])
		if err == io.EOF {
			break
		}
	}
}
