package customer_type

import (
	"testing"
	"fmt"
	"time"
)

type IntConv func(op int) int

//通过自定义类型，让程序有更好的可读性，这里自定义了IntCov类型，这个类型是入参为一个int，返回一个int的函数
func timeSpent(inner func(op int) int) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second *1)
	return op
}

func TestFn(t *testing.T) {
	// 计算函数运行时间
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}