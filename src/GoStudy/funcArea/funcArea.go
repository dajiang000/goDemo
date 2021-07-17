package main

import "fmt"

var str string = "dd"

func demo1() {
	str = "dddddd"
}

func main() {
	fmt.Println(str)
	demo1()
	fmt.Println(str)
}
