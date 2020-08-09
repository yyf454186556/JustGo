package share_mem

import (
	"sync"
	"testing"
)

// //非线程安全的写法
// //(1.00s)
// func TestCounter(t *testing.T) {
	
// 	counter := 0
// 	for i := 0; i < 5000; i++ {
// 		go func() {
// 			counter ++
// 		}()
// 	}

// 	time.Sleep(1 * time.Second)
// 	t.Logf("Counter = %d", counter)
// }

// //线程安全的写法
// // 其实这里也不够好，因为time.Sleep(1 * time.Second)这里是有问题的，如果等的时间太长则效率变低，如果不够长则仍然不能满足要求
// //(1.01s)

// // func TestCoterSafe(t *testing.T) {
// // 	var mut sync.Mutex
// // 	for i := 0; i < 5000; i++ {
// // 		go func() {
// // 			mut.Lock()
// // 			counter ++
// // 			//解锁一般放在defer中，保证一定会执行到
// // 			defer func() {
// // 				mut.Unlock()
// // 			}()
// // 		}()
// // 	}
// // }


// func TestCounterSafe(t *testing.T) {
// 	var mut sync.Mutex
// 	counter := 0
// 	for i := 0; i < 5000; i++ {
// 		go func() {
// 			mut.Lock()
// 			counter ++
// 			//解锁一般放在defer中，保证一定会执行到
// 			defer func() {
// 				mut.Unlock()
// 			}()
// 		}()
// 	}

// 	//注意这里，如果不等待仍然有可能拿不到正确的结果 5000.
// 	//这是因为外面的协程的速度超过了里面的协程的速度
// 	time.Sleep(1 * time.Second)
// 	t.Logf("Counter = %d", counter)
// }

//WaitGroup 只有当wait的内容都执行完成才会继续向下执行
//(0.00s)
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {

		wg.Add(1)
		go func() {
			mut.Lock()
			counter ++
			
			//waitGroup完成
			wg.Done()
			defer func() {
				mut.Unlock()
			}()
		}()
	}
	
	wg.Wait()
	t.Logf("Counter = %d", counter)
}

//以下内容需要自己查找补充一下：
//Go中是支持RW Lock的，即读锁和写锁，由于读锁与读锁之间是不互斥的，这样就会比单纯的互斥锁效率高一些，因此推荐优先使用