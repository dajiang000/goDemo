package main

import (
	"fmt"
)

func putNum(chanInt chan int) {
	for i := 1; i <= 8000; i++ {
		chanInt <- i
	}
	close(chanInt)
}

func GetNum(chanInt chan int, primeInt chan int, chanFlag chan bool) {
	flag := true
	for v := range chanInt {
		flag = true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeInt <- v
		}
	}
	fmt.Println("因为没有数据退出")
	chanFlag <- true
}

func main() {
	chanInt := make(chan int)
	primeInt := make(chan int)
	chanFlag := make(chan bool)

	go putNum(chanInt)

	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go GetNum(chanInt, primeInt, chanFlag)
	go func() {
		for i := 0; i < 16; i++ {
			<-chanFlag
		}
		close(primeInt)
	}()

	for v := range primeInt {
		fmt.Println(v)
	}
}
