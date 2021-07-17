package main

import (
	"fmt"
	"strconv"
)

func main() {

	chanA := make(chan int, 10)
	chanB := make(chan string, 5)
	for i := 0; i < 10; i++ {
		chanA <- i
	}

	for i := 0; i < 5; i++ {
		chanB <- "string " + strconv.Itoa(i)
	}

	for {
		select {
		case v := <-chanA:
			fmt.Println("输出int", v)
		case v := <-chanB:
			fmt.Println("输出string", v)
		default:
			fmt.Println("结束了")
			return
		}
	}
}
