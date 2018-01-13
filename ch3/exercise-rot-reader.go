package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(bs []byte) (n int, err error) {
	n, err = rot.r.Read(bs)
	for i := 0; i < len(bs); i++ {
		if (bs[i] >= 'A' && bs[i] < 'N') || (bs[i] >= 'a' && bs[i] < 'n') {
			bs[i] += 13
		} else if (bs[i] > 'M' && bs[i] <= 'Z') || (bs[i] > 'm' && bs[i] <= 'z') {
			bs[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
