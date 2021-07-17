package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	a := make(chan interface{}, 10)
	a <- Cat{"dayu", 12, 98}
	cat11 := (<-a).(Cat)
	//cat11 := <-a
	fmt.Println(cat11.Name)
}
