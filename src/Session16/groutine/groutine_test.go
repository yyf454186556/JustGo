package groutine_test

import (
	"fmt"
	"time"
	"testing"
)

func TestGroutine(t *testing.T){
	for i:=0;i<10;i++{
		//使用go创建协程，但是需要注意的是：协程函数的j是外部参数i的数据拷贝(Go中都是值传递)
		go func(j int){
			fmt.Println(j)
		}(i)
		//go NT(i)
	}

	time.Sleep(time.Millisecond * 50)
}

func NT(j int){
	fmt.Println(j)
}