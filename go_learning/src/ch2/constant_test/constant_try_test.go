package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry2(t *testing.T) {
	a := 7	// 0111
	//a := 1	// 0001	# 读
	//a := 2	// 0010	# 写
	//a := 4	// 0100	# 可执行
	t.Log(a & Readable == Readable, a & Writeable == Writeable, a & Executable == Executable)
}