package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(1000 * time.Microsecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
