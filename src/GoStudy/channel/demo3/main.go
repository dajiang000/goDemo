package main

import "fmt"

func writeChan(write chan int) {
	for i := 0; i < 50; i++ {
		write <- i
		fmt.Println("写入的数据为", i)
	}
	close(write)
}

func readChan(read chan int, flag chan bool) {
	for v := range read {
		fmt.Println("读取到的结果为", v)
	}
	flag <- true
}

func main() {
	chanIn := make(chan int)
	chanFlag := make(chan bool)
	go writeChan(chanIn)
	go readChan(chanIn, chanFlag)
	fmt.Println(len(chanIn))
	<-chanFlag
}
