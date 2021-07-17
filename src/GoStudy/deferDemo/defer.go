package deferDemo

import "fmt"

func Sum(n *int, m *int) int {
	defer fmt.Println("ok1", *n)
	defer fmt.Println("ok2", *m)
	*n++
	*m++
	return (*n) + (*m)
}
