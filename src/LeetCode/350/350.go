package leetcode_350

func intersect(nums1 []int, nums2 []int) []int {
	myMap := map[int]int{}

	result := []int{}

	for _, num := range nums1 {
		if _,ok := myMap[num]; ok {
			myMap[num] = myMap[num] + 1
		} else {
			myMap[num] = 1
		}
	}

	for _, num := range nums2 {
		if _,ok := myMap[num]; ok && myMap[num] > 0 {
			myMap[num] = myMap[num] - 1
			result = append(result, num)
		}
	}

	return result
}