package dd

import "fmt"

var N int

func init() {
	N = 18
	fmt.Println("dd init()")
}
func GetInt(i int) int {
	return i + 89
}
