package main

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMap(str string) {
	a := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
func main() {
	a := make(map[string]interface{})
	a["名字"] = "安红"
	a["成员"] = []string{"haha", "dd", "mingge"}
	a["队员"] = [3]string{"chuhe"}
	a["尝试"] = map[int]string{1: "dg", 2: "shide"}
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("错误为", err)
	}
	fmt.Println("序列化后的结果", string(data))
	fmt.Printf("反序列化后的结果")
	UnmarshalMap(string(data))
}
