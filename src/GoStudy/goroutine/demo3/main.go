package main

import (
	"fmt"
	"sync"
	"time"
)

var mm = make(map[int]float64)
var lock sync.Mutex

func test(i int) {
	res := 1.0
	for m := 1; m <= i; m++ {
		res *= float64(m)
	}
	lock.Lock()
	mm[i] = res
	lock.Unlock()
}
func main() {
	for i := 1; i <= 100; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	for index, val := range mm {
		fmt.Printf("map[%d]=%f\n", index, val)
	}
}
