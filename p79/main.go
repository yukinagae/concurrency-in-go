package main

import (
	"fmt"
)

func main() {
	var c1, c2 <-chan interface{}
	var c3  chan<- interface{}

	select {
	case <- c1:
		// なにかする
		fmt.Println("c1")
	case <- c2:
		// なにかする
		fmt.Println("c2")
	case c3 <- struct{}{}:
		// なにかする
		fmt.Println("c3")
	}
}
