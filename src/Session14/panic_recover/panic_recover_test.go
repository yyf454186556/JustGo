package panic_recover

import (
	"errors"
	"testing"
	"fmt"
	//"os"
)

func TestPanicVxExit(t *testing.T){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()


	fmt.Println("Start")
	panic(errors.New("Something wrong!")) //会输出调用栈信息，并执行defer
	//os.Exit(-1) //不输出调用栈，不执行defer
}