package demo1

import "testing"

func TestSubDemo(t *testing.T) {
	a := subDemo(2, 4)
	if a != -2 {
		t.Fatalf("有错误，应该是%d,结果是%d", -2, a)
	}
	t.Logf("有错误，应该是%d,结果是%d", -2, a)
}
