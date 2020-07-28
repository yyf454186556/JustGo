package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("otherTask: Now in the otherTask")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("otherTask: Task is done.")
}

//串行写法：
/*
=== RUN   TestService
Done
Now in the otherTask
Task is done.
--- PASS: TestService (0.15s)
*/
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//我们把Channel chan返回，当别的地方需要chan中的结果时，可以等待或直接取。取决于定义chan时有没有声明buffer.
func AsyncService() chan string {
	fmt.Println("AsyncService: AsyncService Start")
	retCh := make(chan string, 1)

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

//不会阻塞的写法：
/*
=== RUN   TestAsyncService
Now in the otherTask
returned result.
Task is done.
Done
--- PASS: TestAsyncService (0.10s)
service exited.
PASS
*/
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

//以下是亲测的重点：
//声明chan的时候有两种方式，retCh := make(chan string) 或者 retCh := make(chan string, 1)
//方式1对应的chan的使用方式需要信息的 制造者 和 消费者 同时都在channel中才会继续执行，如果有一方不在，就会阻塞。
//这种方式下对应的实验结果为
/*
=== RUN   TestAsyncService
AsyncService: AsyncService Start
AsyncService: AsyncService End
otherTask: Now in the otherTask
AsyncService: returned result.
otherTask: Task is done.
Done
--- PASS: TestAsyncService (1.00s)
AsyncService: service exited.
PASS

由结果可知，我们先进入了AsyncService方法，并在此中声明了一个协程来为chan准备消息.
紧接着执行otherTask方法，并在其中等待 1000 ms， 这个时间是远大于 在协程中要执行的 service() 所需要的时间的，但此时没有 消费者 此处会block住

注意这里
AsyncService: AsyncService Start
AsyncService: AsyncService End
otherTask: Now in the otherTask
是同步的，由于准备构造协程需要一定时间，而otherTask中第一行就会输出otherTask: Now in the otherTask，所以导致结果顺序是


AsyncService: AsyncService Start
AsyncService: AsyncService End
otherTask: Now in the otherTask
AsyncService: returned result.

以上是二狗的推测，通过将time.Sleep(time.Millisecond * 1000)移动到otherTask的方法开头验证结果为：
--- PASS: TestService (1.05s)
=== RUN   TestAsyncService
AsyncService: AsyncService Start
AsyncService: AsyncService End
AsyncService: returned result.
otherTask: Now in the otherTask
otherTask: Task is done.
Done
--- PASS: TestAsyncService (1.00s)
我的推断应该是正确的

接着向下，
TestAsyncService在执行otherTask后会继续执行fmt.Println(<-retCh)，这时chan中同时有了生产者和消费者，取消阻塞开始执行。
*/

//对于有缓存的情况，消费者不需要等待了，因此输出结果是：
/*
=== RUN   TestAsyncService
AsyncService: AsyncService Start
AsyncService: AsyncService End
otherTask: Now in the otherTask
AsyncService: returned result.
AsyncService: service exited.
otherTask: Task is done.
Done
--- PASS: TestAsyncService (1.00s)
*/