package leetcode_309

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var dp0 = -prices[0]
	var dp1 = 0
	var dp2 = 1

	for i:= 1; i<len(prices); i++{
		var newdp0 int
		if dp0 > (dp2 - prices[i]){
			newdp0 = dp0
		} else {
			newdp0 = dp2 - prices[i]
		}

		newdp1 := dp0 + prices[i]

		var newdp2 int
		if dp1 > dp2{
			newdp2 = dp1
		} else {
			newdp2 = dp2
		}

		dp0 = newdp0
		dp1 = newdp1
		dp2 = newdp2
	}

	if dp1 > dp2 {
		return dp1
	} else {
		return dp2
	}
}