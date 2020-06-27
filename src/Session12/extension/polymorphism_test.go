package extension

import (
	"testing"
	"fmt"
)

type Code string
type Programmer interface{
	WriteHelloWorld() Code
	//ReadHelloWorld() Code
	//上面这个ReadHelloWorld() Code是一个实验
	/*
		如果一个接口没有实现接口的全部方法会怎样呢？
			如果我们在接口里定义了ReadHelloWorld() Code这个方法，即使后面我们完全没有用到它
			仍然会在build时提示错误：cannot use goProg (type *GoProgrammer) as type Programmer in argument to writeFirstProgram:
				*GoProgrammer does not implement Programmer (missing ReadHelloWorld method)
			
		也就是说，虽然有Duck Type的存在，但如果一个结构没有实现接口的全部方法，那么这个结构就不能作为这个接口的实现。
	*/
}

type GoProgrammer struct{
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct{
}

func (p JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer){
	fmt.Printf("%T, %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T){
	goProg := new(GoProgrammer)
	javaProg := JavaProgrammer{}
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)

	//Interface不只能传递指针类型，而是根据实例有没有实现interface的方法，例如上面的javaProg
	//例如上面如果把 JavaProgram中的WriteHelloWorld方法定义更改为：func (p *JavaProgrammer) WriteHelloWorld() Code
	//则会提示错误：cannot use javaProg (type JavaProgrammer) as type Programmer in argument to writeFirstProgram:
	//JavaProgrammer does not implement Programmer (WriteHelloWorld method has pointer receiver)
}