package main

import "fmt"

type A interface {
	test01()
}
type B interface {
	test02()
}
type C interface {
	test03(int, int) string
	A
	B
}
type D interface {
	test04()
}
type Stu struct {
}

func (stu *Stu) test01() {
	fmt.Println("A")
}
func (stu *Stu) test02() {
	fmt.Println("B")
}
func (stu *Stu) test03(n int, m int) string {
	return "dage"
}
func (stu Stu) test04() {
	fmt.Println("D")
}
func main() {

}
