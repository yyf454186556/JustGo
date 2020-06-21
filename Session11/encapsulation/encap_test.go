package encapsulation

import (
	"testing"
	"fmt"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T){
	e := Employee{"0", "Bob", 20}
    e1 := Employee{Name: "Mike", Age: 30}
    e2 := new(Employee) //注意这里返回的引用/指针, 相当于 e:= &Employee{}
    e2.Id = "2" //与其他主要编程语言的差异：通过实例的指针访问成员不需要使用->
    e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
    //%T 代表输出类型
	t.Logf("e is %T", e) 	//输出 e is encapsulation.Employee
	t.Logf("e2 is %T", e2)  //输出 e2 is *encapsulation.Employee
}

//第一种定义方式在市里对应方法被调用时，实例的成员会进行值复制
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//第二种通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String())
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	/*
		输出为：
		Address is c000064520
			TestStructOperations: encap_test.go:45: ID:0-Name:Bob-Age:20
		Address is c0000644f0
	*/

	t.Log(e.String2())
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	/*
		输出为：
		Address is c0000644f0
			TestStructOperations: encap_test.go:48: ID:0/Name:Bob/Age:20
		Address is c0000644f0
	*/

	//结论，第一种方式（func (e Employee) String() string）会有更大的内存开销，因为是把结构的数据copy了一份，而第二种方式（func (e *Employee) String2() string）引用了相同的地址。
}