package string

import "testing"
import "strings"
import "strconv"

func TestStringFn(t *testing.T){
	s:="A,B,C"
	parts:=strings.Split(s,",")
	for _,item := range parts{
		t.Log(item)
	}
}

func TestConv(t *testing.T){
	//Itoa int转string
	//Atoid string转int
	s:= strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(5+i)
	}
}