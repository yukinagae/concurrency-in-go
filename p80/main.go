package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(1*time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
