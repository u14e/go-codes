package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}


func AsyncService() chan string {
	// 代表带 buffer 的 channel, 容量为 1
	retCh := make(chan string, 1)
	// 不带 buffer 的 channel，需要等待其他程序读取值，才能执行协程里面的下一步
	//retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		// 往 channel 里面放数据
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	// 从 channel 里面取数据
	fmt.Println(<-retCh)
}