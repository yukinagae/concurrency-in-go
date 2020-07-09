package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// メモリ消費量を計測する関数
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	// 絶対に終了しないgoroutine
	noop := func() {
		wg.Done()
		<-c
	}

	// 生成するgoroutineの数
	const numGoroutines = 1e4

	wg.Add(numGoroutines)

	// goroutine生成前のメモリ消費量
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	// goroutine生成後のメモリ消費量
	after := memConsumed()

	// goroutineのメモリ消費量をafter-beforeで計算
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
