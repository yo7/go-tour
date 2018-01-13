package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", s)
}

func Sqrt(x float64) (z float64, err error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z = 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (z * 2))
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
