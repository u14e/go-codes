package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// b = a (不允许隐式类型转换)
	b = int64(a)
	var c MyInt
	// c = b (别名和原有类型也不能进行隐式类型转换)
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 (不支持指针运算)
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)	// 打印类型
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	t.Log(s == "")
}
