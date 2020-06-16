package array_test

import "testing"

func TestArrayInit(t *testing.T){
	var arr [3]int
	t.Log(arr[1],arr[2])

	arr1 := [4]int{1,2,3,4}
	//不指定长度且初始化时可以用[...]
	arr2 := [...]int{1,2,3,4,5}
	arr1[1] = 9
	t.Log(arr1[1], arr1[3])
	t.Log(arr2[2])
}

func TestArrayTraval(t *testing.T){
	arr2 := [...]int{1,2,3,4,5}
	for i:= 0; i < len(arr2); i++{
		t.Log(arr2[i])
	}

	//foreach的写法, idx即为数组的索引的值
	for idx, e:= range arr2 {
		t.Log(idx, e)
	}

	//Go语言中定义的变量一定要使用，所以下面的写法会报错 idx declared but not used
	// for idx, e:= range arr2 {
	// 	t.Log(e)
	// }

	//如果只是想使用foreach的写法，可以使用 _ 来占位
	for _, e:= range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T){
	arr3 := [...] int{1,2,3,4,5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)

	arr3_sec2 := arr3[3:]
	t.Log(arr3_sec2)
}