package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello関数の呼び出し")
}

func main() {
	// 関数をgoroutineで実行
	go hello()

	// 無名関数をgoroutineで実行
	go func() {
		fmt.Println("無名関数の呼び出し")
	}()

	// 無名関数を変数に代入してgoroutineで実行
	foo := func() {
		fmt.Println("無名関数を変数に代入して呼び出し")
	}
	go foo()

	// 1秒sleepして、他のgoroutineが終了するのを待つ
	time.Sleep(1 * time.Second)
}
