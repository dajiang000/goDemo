package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	*A
	*B
}

type D struct {
	a A
}

type F struct {
	Name string
	float64
	n float64
}

func main() {
	c := C{&A{"dajiang"}, &B{"dajiang"}}
	fmt.Println(c)
	f := F{}
	f.float64 = 33.4
	f.n = 98.4
}
