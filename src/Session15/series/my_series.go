package series

import (
	"fmt"
)

func GetFibonacciSerie(n int) []int {
	ret := []int{1,1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2] + ret[i-1])
	}
	return ret
}

//首字母小写在包外不可以访问
func square(n int) int {
	return n * n
}

//这里是在同一个源码文件里的两个同名的init方法，并不会报错，而是会依次执行
func init(){
	fmt.Println("init3")
}

func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init2")
}