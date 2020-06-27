package my_map

import "testing"

func TestInitMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2:=map[int]int{}
	m2[4]=16
	t.Logf("len m2=%d", len(m2))
	m3:=make(map[int]int,10)
	t.Logf("len m3=%d", len(m3))

	//cap()是不能用于求map的capacity的. 下面的code会报错：invalid argument m3 (type map[int]int) for capgo
	//t.Logf("len m3=%d", cap(m3))
}

func TestAccessNotExistingKey(t *testing.T){
	m1:=map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	//以上的输出结果为
	/*
		TestAccessNotExistingKey: map_test.go:21: 0
		TestAccessNotExistingKey: map_test.go:23: 0
	*/
	//  但实际上m1[1]是不存在的值，m1[2]是存在但值为0，这是两种情况 事实上当value不存在时Go会给赋一个默认值
	//好处是这样不会有空指针异常，坏处是需要额外判断一下是不存在还是本身就是零值。具体判断方式如下:

	if v,ok:=m1[3];ok{
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}

	m1[3] = 0

	if v,ok:=m1[3];ok{
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}
}

func TestTraveMap(t *testing.T){
	m1:=map[int]int{1:1,2:4,3:9}
	for key,value := range m1{
		t.Log(key,value)
	}
}