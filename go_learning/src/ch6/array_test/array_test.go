package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr3)
}

func TestArrayLoop(t *testing.T) {
	arr3 := [...]int{3, 4, 5, 6}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	// _ 占位，不想要 index
	for _, e := range arr3 {
		t.Log(e)
	}
}

// 数组截断
func TestArraySection(t *testing.T) {
	arr3 := [...]int{3, 4, 5, 6}
	arr3Sec := arr3[:3]
	t.Log(arr3Sec)
}