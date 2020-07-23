package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1:
		// ここには来ない
	case <-c2:
		// ここには来ない
	default:
		fmt.Printf("In default after %v\n", time.Since(start))
	}
}
