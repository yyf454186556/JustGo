package empty_interface

import (
	"testing"
	"fmt"
)

func DoSomething(p interface{}){
	// if i,ok:=p.(int); ok{
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if i,ok:=p.(string); ok{
	// 	fmt.Println("string", i)
	// 	return
	// }

	// fmt.Println("Unknow Type")

	//使用switch简化上面的结构
	switch v:=p.(type){
		case int:
			fmt.Println("Integer", v)
		case string:
			fmt.Println("string", v)
		default:
			fmt.Println("Unknow Type")
	}

	return
}

func TestEmptyInterfaceAssertion(t *testing.T){
	DoSomething(10)
	DoSomething("15")
	DoSomething('z')
}