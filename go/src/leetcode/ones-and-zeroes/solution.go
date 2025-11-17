// 474. 一和零
// https://leetcode.cn/problems/ones-and-zeroes

package ones_and_zeroes

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for _, str := range strs {
		dp[0], dp[1] = dp[1], dp[0]
		zeros := strings.Count(str, "0")
		ones := len(str) - zeros
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				pick := 0
				if i-zeros >= 0 && j-ones >= 0 {
					pick = dp[0][i-zeros][j-ones] + 1
				}
				dp[1][i][j] = max(dp[0][i][j], pick)
			}
		}
	}
	return dp[1][m][n]
}

func FindMaxForm(strs []string, m int, n int) int {
	return findMaxForm(strs, m, n)
}
