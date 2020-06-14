package operator_test

import "testing"

func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4}
	b := [...]int{1,3,4,5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}

	t.Log(a == b)

	//长度不同不能比较
	//t.Log(a == c) Error: invalid operation: a == c (mismatched types [4]int and [5]int)go
	
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T){
	a:= 7 //0111, (Readable = 1, Writable = 1, Executable = 1)
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
