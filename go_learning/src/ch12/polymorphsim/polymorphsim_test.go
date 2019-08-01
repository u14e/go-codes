package polymorphsim

import (
	"fmt"
	"testing"
)

// 自定义类型
type Code string
// 定义接口
type Programmer interface {
	WriteHelloWorld() Code
}

// 实现接口
type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {

}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

// 多态
func TestPolymorphsim(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}