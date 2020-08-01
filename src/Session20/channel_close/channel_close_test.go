package channel_close

import (
	"testing"
	"fmt"
	"sync"
)


//注释掉的内容是有问题的。
//它利用channel来实现producer和receiver之间的信息传递，但是这种写法receiver需要提前知道channel中可能会传递多少消息，从而进行关闭操作。
//如果利用flag来做标记，比如假定-1就是关闭channel的flag，那么producer同时需要事先知道有多少receiver
//总之就是需要channel的两端能感知对方,没有解耦合。
// func dataProducer(ch chan int, wg *sync.WaitGroup)  {
// 	go func() {
// 		for i:= 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		wg.Done()
// 	}()
// }

// func dataReceiver(ch chan int, wg * sync.WaitGroup) {
// 	go func() {
// 		for i:= 0; i < 10; i++ {
// 			data := <-ch
// 			fmt.Println(data)
// 		}
// 		wg.Done()
// 	}()
// }

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:= 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

		//这行会报错 panic: send on closed channel
		//ch <- 11

		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg * sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok { //ok为bool值,true值表示正常接受,false表示通道关闭
				fmt.Println(data)
			} else {
				break;
			}
		}

		//下面两行的输出结果为0,注意，运行到这里时,ch已经被关闭了，所以这里其实是返回的chan对应的类型的零值.
		// data2 := <-ch
		// fmt.Println(data2)

		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}