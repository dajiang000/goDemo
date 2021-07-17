package main

import (
	"GoStudy/factoryDemo/demo"
	"fmt"
)

func main() {
	stu := demo.NewStu("dashaui", 12)
	fmt.Println(stu)
	fmt.Println(stu.GetScore(), stu.Name)
}
