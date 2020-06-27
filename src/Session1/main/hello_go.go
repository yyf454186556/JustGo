//包，表明代码所在的模块(包，当前在main包)
package main

//引入代码依赖
import (
	"fmt"
	"os"
)

//功能实现
func main() {
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("Command Line args:", os.Args[1])
	}

	fmt.Println("Hello World!")
	os.Exit(-1)
}