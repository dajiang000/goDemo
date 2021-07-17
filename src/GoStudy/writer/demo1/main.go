package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := "E:/一些论文/abc.txt"
	file, err := os.OpenFile(str, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("有错误为", err)
	}

	write := bufio.NewWriter(file) //带缓冲的write
	str1 := "hahahha\n"
	for i := 0; i < 1; i++ {
		write.WriteString(str1)
	}
	write.Flush() //由于带缓冲，所以必须要flush()
	defer file.Close()
}
