package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("E:/一些论文/haha.txt")
	if err != nil {
		fmt.Println("打开文件错误", err)
	}

	defer file.Close()
	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

}
