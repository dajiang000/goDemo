package main

import "fmt"

func fib(n int) []uint64 {
	if n <= 2 {
		if n <= 0 {
			return nil
		} else if n == 1 {
			return []uint64{1}
		} else if n == 2 {
			return []uint64{1, 1}
		}
	}
	fibSlice := make([]uint64, n)
	fibSlice[0] = 1
	fibSlice[1] = 1
	for i := 2; i < n; i++ {
		fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
	}
	return fibSlice
}
func main() {
	var slice1 = make([]int, 2, 4)
	//第一种方式
	slice1 = append(slice1, 5)
	//第二种方式添加，添加的必须是切片，且有...
	slice2 := []int{1, 2, 34, 5}
	slice1 = append(slice1, slice2...)
	fmt.Println("slice1:", slice1)
	slice3 := make([]int, 4)
	slice3[0] = 8
	fmt.Println("原来的slice3:", slice3)
	copy(slice3, slice1)
	fmt.Println("拷贝后的slice3:", slice3)

	str := "蒋小弟"

	slice5 := str[3:] //slice5还是string类型，并不是切片类型，不能使用append和cap
	fmt.Println(slice5)

	str1 := "蒋小哥"
	r := []rune(str1)
	r[1] = '大'
	fmt.Println(string(r))

	fibSlice := fib(12)
	fmt.Println(fibSlice)

	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
}
