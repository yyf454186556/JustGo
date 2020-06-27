package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T){
	m:=map[int]func(op int)int{}
	m[1] =func(op int) int {return op}
	m[2] =func(op int) int {return op * op}
	m[3] =func(op int) int {return op * op * op}

	t.Log(m[1](2),m[2](2),m[3](2))
}

func TestMapForSet(t *testing.T){
	mySet:=map[int]bool{}

	//1) 添加元素
	mySet[1] = true

	n:=1

	//2) 判断元素是否存在
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	//3) 删除元素类似
	delete(mySet, 1)

	//4) 元素个数
	t.Log(len(mySet))


	// 删除元素后显示 1 is not existing
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}