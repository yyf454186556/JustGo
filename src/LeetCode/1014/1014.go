package LeetCode_1014

//https://leetcode-cn.com/problems/best-sightseeing-pair/
//这道题的关键是把得分（A[i] + A[j] + i - j） 拆分成 A[i] + i 和 A[j] - j两部分，然后单次循环时固定A[j] - j

func maxScoreSightseeingPair(A []int) int {
	result := 0
	temp := A[0] + 0

	for j:=1;j<len(A);j++{
		if result < temp + A[j] - j {
			result = temp + A[j] - j
		}

		if temp < A[j] + j {
			temp = A[j] + j
		}
	}

	return result
}