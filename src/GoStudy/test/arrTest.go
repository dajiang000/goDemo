package main

import (
	"fmt"
	"math"
)

func main() {
	var isMin, isMax bool
	score := [8]int{}
	minNum := -1
	maxNum := -1
	maxJudge := 0
	minJudge := 0
	sum := 0
	for i := 0; i < 8; i++ {
		num := 0
		fmt.Printf("输入第%d个成绩", i+1)
		fmt.Scanln(&num)
		score[i] = num
		if minNum == -1 || (minNum > num) {
			minNum = num
			minJudge = i + 1
		}

		if maxNum == -1 || (maxNum < num) {
			maxNum = num
			maxJudge = i + 1
		}
	}

	for _, val := range score {
		if val == minNum && !isMin {
			isMin = true
			continue
		}
		if val == maxNum && !isMax {
			isMax = true
			continue
		}
		sum += val
	}
	average := float64(sum) / 6

	leastDistance := 10000000000.0
	maxDistance := -100000000.9
	worstJudge := 0
	bestJudge := 0
	for i, val := range score {
		if math.Abs(float64(val)-average) > maxDistance {
			maxDistance = float64(val) - average
			worstJudge = i + 1
		}
		if math.Abs(float64(val)-average) < leastDistance {
			leastDistance = float64(val) - average
			bestJudge = i + 1
		}
	}

	fmt.Printf("分最高的裁判：%d,分最低的裁判：%d,平均分：%f，最好裁判：%d,最坏裁判：%d", maxJudge, minJudge, average, bestJudge, worstJudge)
}
