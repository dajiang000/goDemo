package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	content, err := ioutil.ReadFile("E:/一些论文/haha.txt") //一次性读取不带缓冲，只适合用于小的空间
	if err != nil {
		fmt.Println("输出的错误为：", err)
	}
	fmt.Println(string(content)) //content是一个[]byte切片类型

}
