package leetcode_121

import (
	"math"
)

func maxProfit(prices []int) int {
	var result int = 0
	var min int = math.MaxInt64

	for i:=0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] - min > result {
			result = prices[i] - min
		}
	}

	return result
}