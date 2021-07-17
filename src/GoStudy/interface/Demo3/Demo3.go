package main

import "fmt"

type A interface {
	test01()
	test02()
}
type B interface {
	test02()
}
type C interface {
	A
	B
}
type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("A")
}
func (stu Stu) test02() {
	fmt.Println("B")
}
func main() {
	var c C = Stu{}
	c.test02()
}
