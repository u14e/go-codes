// 包，表明代码所在的模块 (包)
package main

// 引入代码依赖
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
	// 返回程序状态
	os.Exit(0)
}