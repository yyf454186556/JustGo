package extension

import (
	"testing"
	"fmt"
)

type Pet struct{
}

func (p *Pet) Speak() {
	fmt.Print("In Pet ...")
}

// Go中不支持重载 以下代码会提示错误：
/*
	method redeclared: Pet.Speak
	method(*Pet) func()
	method(*Pet) func() stringgo
*/
// func (p *Pet) Speak() string {
// 	fmt.Print("In Pet ...")
// 	return "Try to overload"
// }

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("In Pet ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("In Dog ...")
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("In Dog ", host)
	d.p.SpeakTo(host)
}

type Cat struct{
	Pet
}

func (c *Cat) Speak() {
	fmt.Print("Try to overwrite Miao")
}

func TestDog(t *testing.T){
	dog := new(Dog)
	dog.SpeakTo("Hello")

	//注意上面是复合不是集成

	//注意还有Go中的匿名嵌套方法，感觉上像是继承
	/*
		type Cat struct {
			Pet
		}
	*/

	cat := new(Cat)
	cat.SpeakTo("Miao")
	//输出结果是In Pet ...In Pet  Miao(Pet中的方法)

	//可以通过
	var cat1 *Pet = new(Pet)
	//下面一行会提示错误：cannot use new(Cat) (type *Cat) as type *Pet in assignmentgo
	//var cat1 *Pet = new(Cat)
	cat1.Speak()


	//不同的写法也是一样的
	//可以通过
	var cat2 Pet = Pet{}
	//提示错误：cannot use Cat literal (type Cat) as type Pet in assignment
	//var cat2 Pet = Cat{}

	cat2.Speak()
}