package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m["two"])
	t.Logf("len m:%d", len(m))

	m1 := map[string]int{}
	m1["one"] = 1
	t.Logf("len m1:%d", len(m1))

	m2 := make(map[string]int, 10 /* Initial Capacity */)
	t.Logf("len m2:%d", len(m2))
}

func TestAccessNotExistingKet(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])	// key 不存在时，默认是零值
	m1[2] = 0
	t.Log(m1[2])

	// 通过 第二个返回值，判断元素是否存在
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d.", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTraveMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}