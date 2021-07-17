package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFunc(srcName string, aimName string) (written int64, err error) {
	srcFile, err1 := os.Open(srcName)
	if err1 != nil {
		fmt.Println("错误为", err1)
	}
	reader := bufio.NewReader(srcFile)
	aimFile, err2 := os.OpenFile(aimName, os.O_WRONLY|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println("错误为", err2)
	}
	writer := bufio.NewWriter(aimFile)

	return io.Copy(writer, reader)

}
func main() {
	_, err := copyFunc("F:/左神视频/第六章：图算法.mp4", "F:/dajiang.mp4")
	if err != nil {
		fmt.Println("拷贝失败", err)
	} else {
		fmt.Println("拷贝成功")
	}
}
