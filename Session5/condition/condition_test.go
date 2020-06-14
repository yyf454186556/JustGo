package condition_test

import "testing"

func TestIfMultiSec(t *testing.T){
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	/*
		if 支持两段的写法，先给变量赋值，再判断变量的值
		这么写是因为Go中方法支持多个返回值，因此可以有如下的写法

		if v,err := someFunc(); err==nil {
			t.Log("1==1")
		}else{
			t.Log("2==2")
		}
		
	*/
}

func TestSwitchMultiCase(t *testing.T){
	for i:=0;i<5;i++{
		switch i{
			case 0,2:
				t.Log("Even")
			case 1,3:
				t.Log("Odd")
			default:
				t.Log("it is not 0-3")
		}
	}
}

//当作if...else if...else来使用
func TestSwitchCaseCondition(t *testing.T){
	for i:=0;i<5;i++{
		switch {
			case i%2==0:
				t.Log("Even")
			case i%2==1:
				t.Log("Odd")
			default:
				t.Log("Unknow")
		}
	}
}