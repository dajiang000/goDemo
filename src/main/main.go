package main

import (
	"GoStudy/deferDemo"
	"dd"
	"fmt"
	"strings"
)

type myInt int

func kk(a int, b int) {
	if a > 10 {
		return
	}
	fmt.Println("ddd")
}
func getRes(n int) int {
	if n == 1 {
		return 1
	} else {
		return (getRes(n-1) + 1) * 2
	}
}
func demo3(n int, m int) int {
	return n + m
}
func demo4(funcVar func(int, int) int, n int, m int) int {
	return funcVar(n, m)
}
func demo5(n int, m int) (sum int, sub int) {
	sum = n + m
	sub = n - m
	return
}
func demo6(n int, args ...int) int {
	sum := n
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
func init() {
	fmt.Println("main() init")
}
func demo7() func(int) int {
	var i int = 0
	return func(m int) int {
		i = i + m
		return i
	}
}

func demo8(suffix string) func(string) string {
	return func(str string) string {
		if !strings.HasSuffix(str, suffix) {
			return str + suffix
		} else {
			return str
		}
	}
}
func main() {
	fmt.Println(dd.N)
	fmt.Println(demo6(10, 88))
	res := func(n int, m int) int {
		return n - m
	}(12, 13)
	fmt.Println(res)

	f := func(n int, m int) {
		fmt.Println("n*m=", n*m)
	}
	f(1, 4)
	fc := demo7()
	fc1 := demo7()
	fmt.Println(fc(4))
	fmt.Println(fc(4))
	fmt.Println(fc1(4))
	ii := 3
	jj := 4
	fmt.Println(deferDemo.Sum(&ii, &jj))
}
