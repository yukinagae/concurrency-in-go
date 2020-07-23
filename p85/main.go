package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int64
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("Incrementing: %d\n", count)
	}
}
