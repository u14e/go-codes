package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int)  {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func (op int) int) func (op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	// a, _ := returnMultiValues()	// _ 忽略某个返回值
	t.Log(a, b)

	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

// 可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	// 延迟执行函数, 类似 try 中的 finally
	//defer func() {
	//	t.Log("Clear resources")
	//}()
	defer Clear()	// 用于释放资源和锁
	fmt.Println("Started")
	panic("Fatal error")		// 程序抛错，defer依然执行
}