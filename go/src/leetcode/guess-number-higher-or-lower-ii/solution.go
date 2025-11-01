// 375. 猜数字大小 II
// https://leetcode.cn/problems/guess-number-higher-or-lower-ii

package guess_number_higher_or_lower_ii

import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for l := 1; l < n; l++ {
		for i := 1; i <= n-l; i++ {
			j := i + l
			v := math.MaxInt
			for k := i; k < j; k++ {
				v = min(v, max(dp[i][k-1], dp[k+1][j])+k)
			}
			dp[i][j] = v
		}
	}
	return dp[1][n]
}

func GetMoneyAmount(n int) int {
	return getMoneyAmount(n)
}
