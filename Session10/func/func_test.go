package func_test

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

// 函数多个返回值
func returnMultiValues() (int, int) {
	return  rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int{
	time.Sleep(time.Second *1)
	fmt.Println("In SlowFun, the parameter is", op)
	return op
}


// 传递一个函数进来
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}


func TestFn(t *testing.T) {
	// 函数多个返回值
	a, b := returnMultiValues()
	t.Log(a, b)

	// 计算函数运行时间
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}