package client

import (
	"fmt"
	"testing"
	"Session15/series"
)

//import "Session15/series"

func TestPackage(t *testing.T){
	fmt.Println(series.GetFibonacciSerie(5))
}