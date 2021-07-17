package main

import "fmt"

func main() {
	a := make(chan int, 3)
	a <- 3
	close(a)
	//a<-4
	b := <-a
	fmt.Println(b)
}
