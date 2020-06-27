package client

import (
	"testing"
	"Session15/series"
)

func TestPackage(t *testing.T){
	t.Log(series.GetFibonacciSerie(5))
	//t.Log(series.square(5)) //包中小写的方法是不能被访问的，会提示错误cannot refer to unexported name series.squareg
}