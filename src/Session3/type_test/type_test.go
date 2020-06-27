package type_test

import "testing"

type MyInt int64

func  TestImplicit(t *testing.T)  {
	var a int = 1
	var b int64

	//不支持隐式类型转换
	//b = a //Error: cannot use a (type int) as type int64 in assignmentgo
	b = int64(a)
	t.Log(a,b)

	//别名也不行
	var c MyInt
	//c = b //Error: cannot use b (type int64) as type MyInt in assignmentgo

	c = MyInt(b)
	t.Log(a,b,c)
}

func TestPoint(t * testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 输出类型，使用 "%T"

	//不支持指针运算
	//aPtr = aPtr + 1 Error: invalid operation: aPtr + 1 (mismatched types *int and int)
}

func TestString(t *testing.T) {
	var s string  //在Go中是数值类型，注意和其他语言的区别
	t.Log("*" + s + "*")   //output: **
	t.Log(len(s))          //output: 0

	//在Golang中 string 的初始值是 空，因此相关判断时要注意 if(s == "")
}