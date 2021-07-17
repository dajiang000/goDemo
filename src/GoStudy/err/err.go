package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 0
	num2 := 100
	fmt.Println(num2 / num1)
}

func main() {
	test()
	fmt.Println("哈哈，活着到这里了")
}
