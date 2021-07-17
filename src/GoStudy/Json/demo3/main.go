package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []interface{}
	a = append(a, map[string]string{"dd": "dagad", "dg": "4444"}, 444, []float64{3.456, 4.56})
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("错误为", err)
	}
	fmt.Println(string(data))
}
