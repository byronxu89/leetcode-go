package algo6

import "math"

func minimumTotal(triangle [][]int) int {

	/*
	   [2]
	   [3],[4]
	   [6],[5],[7]
	*/
	rows := len(triangle)
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]

	}

	ans := math.MaxInt32
	lastrow := rows - 1
	for i := 0; i < rows; i++ {
		ans = min(ans, dp[lastrow][i])
	}

	return ans

}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
