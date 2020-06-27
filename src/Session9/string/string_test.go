package string_test

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	t.Log(len(s))

	s = "\xE4\xB8\xA4" //即使是不可显示的字符也可以存储，如"\xE4\xBB\xA5"将显示乱

	t.Log(s)
	t.Log(len(s)) //输出结果为3，但其实对应的是一个中文字符 ‘两’

	// string是不可以变的byte slice，因此不能给它赋值
	a := "Hello"
	//a[1] = '3' // 报错：cannot assign to a[1]go
	t.Log(a)

	b := "中"
	t.Log(len(b)) //是byte数
	c := []rune(b) //rune能够取出字符串中的unicode,这是Go的一个内置机制
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
	//以上四个log输出的结果分别为：
	/*
	    TestString: string_test.go:21: 3
		TestString: string_test.go:23: 1
		TestString: string_test.go:24: 中 unicode 4e2d
		TestString: string_test.go:25: 中 UTF8 e4b8a4
	*/
}

func TestStringToRune(t *testing.T){
	s:= "中华人民共和国"
	for  _,c:=range s{
		t.Logf("%[1]c %[1]d",c)
	}

	/*d
		以上运行结果为
		    TestStringToRune: string_test.go:38: 中 20013
			TestStringToRune: string_test.go:38: 华 21326
			TestStringToRune: string_test.go:38: 人 20154
			TestStringToRune: string_test.go:38: 民 27665
			TestStringToRune: string_test.go:38: 共 20849
			TestStringToRune: string_test.go:38: 和 21644
			TestStringToRune: string_test.go:38: 国 22269

		注意这里看代码可能有些别扭，这里的‘%[1]c %[1]d’ 其实意思是 以%c的格式格式化第1个字符 和 以%d的格式格式化第1个字符
	*/
}