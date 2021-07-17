package demo2

import (
	"fmt"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := &Monster{"dajiang", 23, "nan", 78.9, "coding"}
	res := monster.Store("F:/abc.txt")
	if !res {
		t.Fatal("写入不成功")
	}
	t.Log("写入成功")
}
func TestMonster_Restore(t *testing.T) {
	monster := &Monster{}
	res := monster.Restore("F:/abc.txt")
	fmt.Println(monster)
	if !res || monster.Name != "dajiang" {
		t.Fatal("读取不成功")
	}
	t.Log("读取成功")
}
