package main

import (
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
}
