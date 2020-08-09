package util_anyone_reply

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


//普通channel, 会造成内存泄漏
// func FirstResposne() string {
// 	numOfRunner := 10
// 	ch := make(chan string)

// 	for i:=0; i < numOfRunner; i++{
// 		go func(i int) {
// 			ret := runTask(i)
// 			ch <- ret
// 		}(i)
// 	}

// 	return <-ch
// }

//普通channel
func FirstResposne() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i:=0; i < numOfRunner; i++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

//以下是其中一组测试结果，可以看到runTask方法确实只执行了一次，等到第一个response后就结束了.ch关闭，但上面的写法是有严重的问题的
//但上面的写法是有严重的问题的
//但上面的写法是有严重的问题的
//但上面的写法是有严重的问题的
//重要的事情再说三遍,因为使用的是没有使用buffer的chan,结束后会有一些协程没有被关闭，大量运行会导致巨大的开销，内存泄漏
/*
=== RUN   TestFirstResponse
    TestFirstResponse: first_response_test.go:38: Before: 2
    TestFirstResponse: first_response_test.go:39: The result is from 8
    TestFirstResponse: first_response_test.go:41: After: 11
--- PASS: TestFirstResponse (1.01s)
PASS
*/
func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResposne())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

//使用了BufferChannel的运行结果如下
/*

=== RUN   TestFirstResponse
    TestFirstResponse: first_response_test.go:62: Before: 2
    TestFirstResponse: first_response_test.go:63: The result is from 3
    TestFirstResponse: first_response_test.go:65: After: 2
--- PASS: TestFirstResponse (1.01s)
PASS

*/

//再复习一遍 Channel和BufferChannel
/*
	与普通的Channel不同, BufferChannel每一个消息的产生这往里面放消息，只要是有buffer的，都不需要等待接收者把消息拿走,它们之间是解耦和的关系.(消息的生产者和消费者之间)
	对应到上面的代码,无Buffer的情况下,每个协程调用的runTask都向chan中放了消息,但FirstResposne只接收了第一个,其他九个协程只好等待，所以不能被关闭.
	而Bufferchan则是生产者向里丢消息后就可以之间关闭协程.
*/