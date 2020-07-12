package leetcode_122

//题目要求其实可以简化为在一个数组中找到两个索引a, b (a < b) 使得两者的差值最大.
//这么一想这个思路好像有点意思，只要遇到后面的数值大于前面的，就求差值并入结果中。
//不论中间是连续递增还是不连续的，其实并不影响两端的差值

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var result int = 0
	
	for i:=1; i < len(prices); i++ {
		if prices[i] - prices[i-1] > 0 {
			result += prices[i] - prices[i-1]
		}
	}

	return result
}