package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("分线程+", i)
		time.Sleep(time.Second)
	}
}
func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("主线程+", i)
		time.Sleep(3 * time.Second)
	}
}
