package slice_test

import "testing"

func TestSliceInit(t *testing.T){
	//与声明数组很类似，区别是没有指定长度！
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	//初始化
	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	//这里s2长度为3，但容量为5
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//这里会报错，因为后两个元素会越界
	//t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	t.Log(s2[0],s2[1],s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0],s2[1],s2[2],s2[3])
}

func TestSliceGrowing(t *testing.T){
	s:=[]int{}
	for i:=0; i<10; i++{
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	//输出结果
	/*
	    TestSliceGrowing: slice_test.go:30: 1 1
		TestSliceGrowing: slice_test.go:30: 2 2
		TestSliceGrowing: slice_test.go:30: 3 4
		TestSliceGrowing: slice_test.go:30: 4 4
		TestSliceGrowing: slice_test.go:30: 5 8
		TestSliceGrowing: slice_test.go:30: 6 8
		TestSliceGrowing: slice_test.go:30: 7 8
		TestSliceGrowing: slice_test.go:30: 8 8
		TestSliceGrowing: slice_test.go:30: 9 16
		TestSliceGrowing: slice_test.go:30: 10 16
	*/

	//每次不够放的时候，都会将上次的扩展为之前的两倍，因此要使用类似于s2 = append(s2, 1) 因为是新的空间给s2，并将原来的数据拷贝过去。
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	//输出结果为
	//TestSliceShareMemory: slice_test.go:53: [Apr May Jun] 3 9,容量是直接到末尾的，所以是9.

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unkonw"
	t.Log(summer, len(summer), cap(summer))
	t.Log(q2, len(q2), cap(q2))
	//输出结果为
	/*
	    TestSliceShareMemory: slice_test.go:60: [Unkonw Jul Aug] 3 7
    	TestSliceShareMemory: slice_test.go:61: [Apr May Unkonw] 3 9
	*/

	//由于是共享的存储空间，因此summer的改动同时影响到了q2的值.
}

func TestSliceComparing(t *testing.T){
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}

	//会报错，nvalid operation: a == b (slice can only be compared to nil)
	// if(a == b){
	// 	t.Log("Can Compare")
	// }

	t.Log(a[2])
	t.Log(b[2])
	b[2] = 9
	t.Log(a[2])
	t.Log(b[2])

	
}