package try_test

import "testing"

func TestFibList(t *testing.T){
	//Go支持多种复制关系以及一定的类型推断，但在2.0版本后将会不推荐甚至禁用一部分写法，目的是code的风格统一
	var a int = 1
	var b int = 1
	/*
		var (
			a int = 1
			b = 1
		)
		// a:=1
		// b:=1
	*/
	for i:=0;i<5;i++{
		t.Log(" ", b)
		tmp:=a
		a=b
		b=tmp + a
	}
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T){
	t.Log(Monday, Wednesday)
} 