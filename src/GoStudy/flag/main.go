package main

import (
	"flag"
)

func main() {
	var user string
	var port string
	var age int

	flag.StringVar(&user, "user", "", "用户名为")
	flag.StringVar(&port, "port", "", "端口为")
	flag.IntVar(&age, "age", 0, "年龄为")
	flag.Parse()

}
