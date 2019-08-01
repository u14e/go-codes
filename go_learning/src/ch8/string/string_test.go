package string_test

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log(s)						// 初始化为默认零值
	s = "hello"
	// s[1] = '3'				// string 是不可变的 byte slice
	t.Log(len(s))

	s = "\xE4\xB8\xA5"	// 可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s))				// len 函数返回它所包含的 byte 数

	s = "中"
	t.Log(len(s))				// 3 是 byte 数

	c := []rune(s)			// 取出 Unicode
	t.Log(c)
	t.Logf("中 Unicode: %X", c[0])	// 十六进制表示
	t.Logf("中 UTF8: %X", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}