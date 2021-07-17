package main

import "fmt"

type A struct {
	x int
	y int
}
type B struct {
	x string
	y int
}

func main() {
	var aa interface{}
	bb := A{1, 3}
	aa = bb
	cc, ok := aa.(A)
	if ok {
		fmt.Println("aa是A类型")
	} else {
		fmt.Println("aa不是A类型")
	}
	fmt.Println(cc)

}
