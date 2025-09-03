// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square

package maximal_square

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	s := 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			s = max(s, dp[i][0])
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			s = max(s, dp[0][j])
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				s = max(s, dp[i][j])
			}
		}
	}
	return s * s
}

func MaximalSquare(matrix [][]byte) int {
	return maximalSquare(matrix)
}
