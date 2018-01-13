package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (z * 2))
	}
	return
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
