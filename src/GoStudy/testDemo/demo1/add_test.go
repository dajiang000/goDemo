package demo1

import (
	"testing"
)

func TestAddDemo(t *testing.T) {
	a := addDemo(1, 3)
	if a != 4 {
		t.Fatalf("输出的数为%d,实际为%d", a, 4)
	}
	t.Logf("测试成功")
}

//func TestSubDemo(t *testing.T){
//	a:=subDemo(2,4)
//	if a != -2{
//		t.Fatalf("有错误，应该是%d,结果是%d",-2,a)
//	}
//	t.Log("结果正确")
//}
