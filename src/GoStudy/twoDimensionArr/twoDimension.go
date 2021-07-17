package main

import "fmt"

func main() {
	var n [3][4]int = [...][4]int{{1, 2, 3}, {1, 2, 4, 2}, {}}
	fmt.Println(n)
	//for i:=0;i<len(n);i++{
	//	for j:=0;j<4;j++{
	//		fmt.Printf("  %d",n[i][j])
	//	}
	//}

	for _, val := range n {
		for _, val2 := range val {
			fmt.Printf("  %d", val2)
		}
		fmt.Println()
	}

}
