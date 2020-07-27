package leetcode_120

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	return back(0,0,triangle)
}

func back(i int, j int, triangle [][]int) int {
	if i >= len(triangle) {
		return 0
	}

	var x = back(i+1, j, triangle)
	var y = back(i+1, j+1, triangle)

	if x > y {
		return triangle[i][j] + y
	} else {
		return triangle[i][j] + x
	}
}