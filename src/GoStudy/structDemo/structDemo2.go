package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Num  int
}
type B struct {
	Name string
	Num  int
}

type Monster struct {
	Name  string `json:"name"` //tag 标记json返回时，将Name转换为name
	Age   int    `json:"age"`
	Skill string `json:"skill"`
	//Name string
	//Age int
	//Skill string
}
type integer int

type Student struct {
	Name string
	Age  int
}

type MethodUtils struct {
}

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) ticket() {
	if visitor.Age < 18 && visitor.Age >= 0 {
		fmt.Printf("%s,太小了，回家玩吧\n", visitor.Name)
	} else if visitor.Age >= 18 && visitor.Age < 65 {
		fmt.Printf("%s,你的票价为20元\n", visitor.Name)
	} else if visitor.Age >= 65 && visitor.Age <= 100 {
		fmt.Printf("%s,免费", visitor.Name)
	}
}
func (Mu *MethodUtils) Print(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
func (Mu *MethodUtils) Print2(n [][]int) {
	len := len(n)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			n[i][j], n[j][i] = n[j][i], n[i][j]
		}
	}
}

func (student *Student) String() string {
	str := fmt.Sprintf("学生信息%s,%d", student.Name, student.Age)
	return str
}
func (i integer) print() {
	fmt.Println(i)
}

func (monster *Monster) ChangeName(name string) string {
	monster.Name = name
	return name
}
func main() {
	a := A{}
	b := B{}

	a = A(b)
	fmt.Println(a, b)

	monster := Monster{"白骨精", 1000, "变美女"}

	jsonStr, err := json.Marshal(monster)
	if err == nil {
		fmt.Printf("输出的字符串%s\n", string(jsonStr))
	}

	name := monster.ChangeName("蜘蛛精")
	fmt.Println(monster, name)

	i := integer(1)
	i.print()
	student := Student{"dajiang", 125}
	fmt.Println(&student)

	mu := MethodUtils{}
	mu.Print(6)

	n := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mu.Print2(n)
	for _, val := range n {
		for _, val1 := range val {
			fmt.Print(val1, "\t")
		}
		fmt.Println()
	}
	var visitor Visitor
	for {
		fmt.Println("输入你的名字")
		fmt.Scanln(&visitor.Name)
		if visitor.Name == "n" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("输入年龄")
		fmt.Scanln(&visitor.Age)
		visitor.ticket()
	}
}
