package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string  `json:"mingzi"`
	Age   int     `json:"nianling"`
	Score float64 `json:"chenji"`
}

func UnmarshalStruct(a string) {
	stu := Student{}
	err := json.Unmarshal([]byte(a), &stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}
func main() {
	student := Student{"dajiang", 34, 76.98}

	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("错误为", err)
	}
	//fmt.Println(string(data))
	UnmarshalStruct(string(data))
}
