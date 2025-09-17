// 279. 完全平方数
// https://leetcode.cn/problems/perfect-squares/

package perfect_squares

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; ; i++ {
		j := i * i
		if j > n {
			break
		}
		dp[j] = 1
	}

	for i := 1; i <= n; i++ {
		if dp[i] > 0 {
			continue
		}
		mid := i / 2
		x := math.MaxInt
		for j := 1; j <= mid; j++ {
			x = min(x, dp[j]+dp[i-j])
		}
		dp[i] = x
	}
	return dp[n]
}

func NumSquares(n int) int {
	return numSquares(n)
}
