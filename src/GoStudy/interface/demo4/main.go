package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	Name  string
	Score int
	Age   int
}

type stuSlice []Student

func (slice stuSlice) Len() int {
	return len(slice)
}
func (slice stuSlice) Less(i int, j int) bool {
	return slice[i].Score < slice[j].Score
}
func (slice stuSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var stu stuSlice = stuSlice{}
	for i := 0; i < 10; i++ {
		temp := Student{}
		temp.Name = fmt.Sprint("这是第", i+1, "号")
		temp.Age = rand.Intn(100)
		temp.Score = rand.Intn(80)
		stu = append(stu, temp)
	}
	for _, val := range stu {
		fmt.Println("姓名：", val.Name, "年龄：", val.Age, "成绩：", val.Score)
	}
	sort.Sort(stu)
	fmt.Println("------------------------------------------------")
	for _, val := range stu {
		fmt.Println("姓名：", val.Name, "年龄：", val.Age, "成绩：", val.Score)
	}
}
