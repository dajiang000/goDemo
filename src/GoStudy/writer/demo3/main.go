package main

import (
	"fmt"
	"os"
)

func main() {
	file1Path := "E:/一些论文/abc.txt"
	file2Path := "F:/dajiang.txt"

	content, err := os.ReadFile(file1Path)
	if err != nil {
		fmt.Println("有错误", err)
	}

	err = os.WriteFile(file2Path, content, 066)
	if err != nil {
		fmt.Println("有错误", err)
	}
}
