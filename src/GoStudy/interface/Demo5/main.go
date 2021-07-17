package main

import (
	"fmt"
)

type A interface {
	test()
}
type s1 struct {
	Name string
}
type s2 struct {
	Name string
}

func (S1 s1) test() {
	fmt.Println(S1.Name)
}
func (S2 s2) test() {
	fmt.Println(S2.Name)
}
func change(c [3]A) {
	c[0] = s1{"1111"}
}
func main() {
	c := [3]A{}
	c[0] = s1{"dage"}
	c[1] = s2{"hehe"}
	c[2] = s1{"dajiang"}
	fmt.Println(c)
	change(c)
	fmt.Println(c)
}
