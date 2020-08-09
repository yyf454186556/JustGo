package singleton_test

import (
	"unsafe"
	"testing"
	"fmt"
	"sync"
)

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func(){
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

//这里会输出一下十个协程中创建的对象的地址,以下是运行结果
/*
Create Obj
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
0x6873b8
*/
//对象只被创建了一次
func TestTGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	for i:= 0; i<10; i++ {
		wg.Add(1)
		go func(){
			obj := GetSingletonObj()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}