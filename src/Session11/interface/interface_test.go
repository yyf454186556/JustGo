package interface_test

import (
	"testing"
)

type Programmer interface{
	WriteHelloWorld() string
}

type GoProgrammer struct{

}

func (g *GoProgrammer) WriteHelloWorld() string{
	return "fmt.Println(\"Hello World!\")"
}

/*注意这里，其实p定义的是接口Programmer，而p = new(GoProgrammer)
  是将接口Programmer的具体实现'GoProgrammer'作为p的实例
  这里没有使用传统的【接口定义，接口实现继承自接口定义，具体使用的地方利用接口定义通过容器或别的方式获取接口实现】的传统方法
  而是使用了DuckType.所以DuckType就是指，这个鸟虽然我不知道是什么，但是看起来脚上有蹼，扁嘴，像是鸭子，那么将当它是鸭子。(2333,这是老师的原话)
  对应上面就是接口GoProgrammer所对应的方法‘WriteHelloWorld’与Programmer中定义的方法看起来是一样的，那么我们就当GoProgrammer是Programmer的具体实现
*/

func TestClinet(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}