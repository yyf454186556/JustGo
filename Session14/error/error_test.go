package error_test

import (
	"fmt"
	"errors"
	"testing"
)

//区分错误类型，定义预置错误
var LessThanTwoError = errors.New("n should be not less than 2")
var LagerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([] int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LagerThanHundredError
	}


	fibList := []int{1,1}

	for i:=2; i<n; i++{
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T){
	if v,err:=GetFibonacci(-10);err != nil{
		if err == LessThanTwoError {
			fmt.Println("Less than 2")
		} else if err == LagerThanHundredError {
			fmt.Println("Larger than 100")
		}
	}else{
		t.Log(v)
	}
}
  