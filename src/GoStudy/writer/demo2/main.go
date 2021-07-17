package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	name := "E:/一些论文/abc.txt"
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("有错误，", err)
	}

	reader := bufio.NewReader(file)
	for {
		str2, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str2)
	}
	str2 := "hahahha"

	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str2)
	}
	writer.Flush()
	defer file.Close()
}
