package cancel_by_close

import (
	"time"
	"fmt"
	"testing"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <- cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}


func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)

	//开启五个协程
	for i:= 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancel Done")
		}(i, cancelChan)
	}

	//cancel_1只是随便传了个标记关闭的值,这样在isCancelled的方法中的channel收到消息，知道应该取消任务了，于是返回了true.但这种情况下，五个协程里只有一个收到了消息，被取消掉了.
	//因为五个协程都在等待chan中的消息，而只有一个收到了
	//cancel_1(cancelChan)

	//利用channel的关闭机制(关闭时会向所有的receiver发消息，提示关闭了，如果没有设置判断是否关闭的参数，关闭时会立即返回channel对应的类型的零值)，五个协程都被正确关闭了.
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}