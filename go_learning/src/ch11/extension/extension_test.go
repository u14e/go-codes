package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//	//d.p.SpeakTo(host)
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()								// Wang!
	// 父类的 SpeakTo 里面没有调用子类重写的 Speak
	dog.SpeakTo("u14e")	// ...  u14e
}