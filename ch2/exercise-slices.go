package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (s [][]uint8) {
	s = make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
	}

	for i := range s {
		for j := range s[i] {
			switch {
			case j%15 == 0:
				s[i][j] = 240
			case j%3 == 0:
				s[i][j] = 120
			case j%5 == 0:
				s[i][j] = 150
			default:
				s[i][j] = 100
			}
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}
