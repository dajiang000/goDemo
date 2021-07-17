package main

import "fmt"

func test2(arr *[3]int) {
	(*arr)[0] = 11
}
func main() {
	var num1 [3]int = [3]int{1, 2, 4}
	fmt.Println(num1)

	var num2 = [3]int{4, 5, 6}
	fmt.Println(num2)

	var num3 = [...]int{9, 10, 11}
	fmt.Println(num3)

	num4 := [...]int{12, 13, 14}
	fmt.Println(num4)

	num5 := [...]int{22, 33, 44, 55}
	for index, value := range num5 {
		fmt.Printf("输出的下标%d,输出的结果%d\n", index+1, value)
	}

	num6 := [3]int{1, 2, 3}
	fmt.Println(num6)
	test2(&num6)
	fmt.Println(num6)

	num7 := [...]int{11, 22, 33, 44, 55}
	slice1 := num7[1:3] // 范围为1-3 但是不包括3
	fmt.Printf("类型:%T\n", slice1)
	fmt.Println("元素是：", slice1) //输出元素
	fmt.Println(cap(slice1))    //输出当前容量
	fmt.Println(slice1[0])      //提取slice中的值
}
