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

func Sum (ops ... int) int {
	ret := 0
	for _, s := range ops {
		ret += s
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(5,6,7,8,9))
}

func Clear() {
	fmt.Println("Clear resources in Clear function.")
}

func TestDefer(t *testing.T){
	defer func() { //这个内联函数直到TestDefer执行完毕，返回前才会执行。 通常用于清理资源，释放锁等功能
		defer Clear()
		t.Log("Clear resources")
	}()
	t.Log("Started")
	panic("Fatal error") //defer 仍会执行 //panic是Go中程序异常中断，抛出致命错误
}