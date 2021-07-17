package main

import (
	"fmt"
	"strconv"
	"time"
)

func test4() {
	c := 100
	str := ""
	for j := 0; j < 1; j++ {
		str += strconv.Itoa(c)
	}
}
func main() {
	now := time.Now()
	fmt.Println(now.YearDay())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(time.Nanosecond)

	now1 := time.Now()
	fmt.Println(now1.UnixNano())
	fmt.Println(now1.Unix())

	/*
		判断一个函数执行多长时间
	*/

	nowTime := now.UnixNano()
	test4()
	nowTime1 := now.UnixNano()
	fmt.Println(nowTime1 - nowTime)

	n := 1
	n1 := &n
	n2 := &n1
	n3 := &n2
	n4 := new(int)
	*n4 = 100
	fmt.Printf("%T,%T,%T,%v,%v,%v", n1, n2, n3, n1, n2, n3)
}
