package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/一些论文/haha.txt")
	if err != nil {
		fmt.Println("打开文件有误,问题为：", err)
	}
	fmt.Println(file.Name())

	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件出错，问题为：", err)
	}
}
