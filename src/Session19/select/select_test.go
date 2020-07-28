package select_test

import (
	"time"
	"testing"
	"fmt"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	fmt.Println("AsyncService: AsyncService Start")
	retCh := make(chan string)

	//这里是为了当AsyncService被调用时启动另一条协程去运行而不阻塞当前的协程
	go func() {
		ret := service()
		fmt.Println("AsyncService: returned result.")
		retCh <- ret
		fmt.Println("AsyncService: service exited.")
	}()

	fmt.Println("AsyncService: AsyncService End")
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <- AsyncService():
			fmt.Println("Success " + ret)
	case <- time.After(time.Millisecond * 200):
			fmt.Println("Time out")
	}
}