package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	/*=== num 1 ===*/
	//var a int = 1
	//var b int = 1

	/*=== num 2 ===*/
	//var (
	//	a int = 1
	//	b = 1	// 自动推断 变量类型
	//)

	/*=== num 3 (推荐) ===*/
	a := 1
	b := 1

	/*=== num 4 ===*/
	// 先声明，后赋值
	//var a int
	//a = 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		a, b = b, a + b
	}
}
