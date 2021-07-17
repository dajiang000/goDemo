package main

import "fmt"

type Stu struct {
}
type interfaceA interface {
	Say()
}
type integer int

func (i integer) Say() {
	fmt.Println("integer :", i)
}
func (stu *Stu) Say() {
	fmt.Println("stu say")
}

func main() {
	stu := &Stu{}
	var interfaceC interfaceA = stu
	interfaceC.Say()

	var i integer = 10
	var interfaceB interfaceA = i
	interfaceB.Say()
}
