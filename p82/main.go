package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	select {
	case <-c:
		// ここは通らない
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
