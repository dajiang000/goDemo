package main

import (
	"fmt"
)

type Cat struct {
	Age    int
	Gender string
	Color  string
}

func f() {
	cat1 := Cat{Gender: "mu", Age: 12, Color: "yellow"}
	cat1.Color = "red"
	fmt.Println(cat1)
	c := [3]int{}
	c[0] = 1
	c[2] = 4
	fmt.Println(c)
	fmt.Println("---------------------------------------------------")
	cat3 := Cat{Gender: "xixi", Age: 12, Color: "yellow"}
	fmt.Println(cat3)
	cat4 := &cat3
	cat4.Gender = "gong"
	fmt.Println(cat3)
	fmt.Println(cat4)

	var cat5 Cat
	cat5.Gender = "nu"
	cat5.Color = "blue"
	cat5.Age = 34
	cat6 := new(Cat)
	cat6.Gender = "nan"
	cat6.Color = "black"
	cat6.Age = 2
	fmt.Println(cat6)
	cat9 := cat6
	cat9.Age = 55
	fmt.Println(cat6)
	cat7 := &Cat{11, "nan", "yellow"}
	fmt.Println(cat7)
	cat8 := cat7
	cat8.Age = 77
	fmt.Println(cat7)
}
