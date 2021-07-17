package main

import (
	"fmt"
	"sort"
)

func main() {
	a := map[string]int{
		"dajiang": 12,
		"haha":    24,
	}
	fmt.Println(a)

	b := make(map[string]string)
	b["dage"] = "dd"
	b["ddd"] = "gdaf"
	val, res := b["dd"]
	if res {
		fmt.Println("找到了", val)
	} else {
		fmt.Printf("没找到%v\n", val)
	}

	for key, value := range b {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}

	student := make(map[string]map[string]string)
	student["dage"] = make(map[string]string)
	student["dage"]["name"] = "阿绵"
	student["dage"]["score"] = "112"
	student["dage"]["gender"] = "男"

	student["dajie"] = make(map[string]string)
	student["dajie"]["name"] = "huadajie"
	student["dajie"]["score"] = "140"
	student["dajie"]["gender"] = "女"

	for k1, val1 := range student {
		fmt.Printf("k1:%s\n", k1)
		for k2, val2 := range val1 {
			fmt.Printf("key:%s , value:%s\n", k2, val2)
		}
	}

	mapSlice := make([]map[string]string, 2)
	mapSlice[0] = make(map[string]string)
	mapSlice[0]["name"] = "dachen"
	mapSlice[0]["score"] = "56"
	mapSlice[1] = make(map[string]string)
	mapSlice[1]["name"] = "dajiang"
	mapSlice[1]["score"] = "44"
	mapSlice[1]["gender"] = "男"

	newMap := make(map[string]string)
	newMap["name"] = "阿绵"
	newMap["gender"] = "女"
	mapSlice = append(mapSlice, newMap)
	for _, val := range mapSlice {
		for key, val1 := range val {
			fmt.Printf("key:%s,value:%s\n", key, val1)
		}
	}

	score := make(map[int]int)
	score[14] = 45
	score[1] = 56
	score[89] = 11
	score[991] = 194
	slice1 := make([]int, 0)
	for key, _ := range score {
		slice1 = append(slice1, key)
	}

	sort.Ints(slice1)
	for _, value := range slice1 {
		fmt.Printf("key:%d,value:%d\n", value, score[value])
	}
}
