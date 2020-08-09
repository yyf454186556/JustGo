package until_all_done

import (
	"testing"
	"time"
	"runtime"
	"fmt"
)



func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResposne() string {
	numOfRunner := 10
	ch := make(chan string)

	for i:=0; i < numOfRunner; i++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""

	for j:=0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

//运行结果如下:
/*
=== RUN   TestAllResponse
    TestAllResponse: until_all_done_test.go:39: Before: 2
    TestAllResponse: until_all_done_test.go:40: The result is from 8
        The result is from 9
        The result is from 7
        The result is from 6
        The result is from 0
        The result is from 1
        The result is from 4
        The result is from 2
        The result is from 5
        The result is from 3
        
    TestAllResponse: until_all_done_test.go:42: After: 2
--- PASS: TestAllResponse (1.01s)
PASS
*/

//当然也可以用WaitGroup来实现了
func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResposne())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
