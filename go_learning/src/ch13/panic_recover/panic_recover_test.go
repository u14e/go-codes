package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally!")
	//}()

	// 恢复错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
		fmt.Println("Finally!")
	}()

	fmt.Println("Start")

	panic(errors.New("Something Wrong!"))
	//os.Exit(-1)
}


