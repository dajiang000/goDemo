package main

import (
	"GoStudy/fengzhuangDemo/demo"
	"fmt"
)

func main() {
	a := demo.NewAccount("大蒋")
	a.SetPwd("dashao")
	a.SetBalance(45)
	fmt.Println(a)
	fmt.Println(a.GetBalance(), a.GetPwd())
}
