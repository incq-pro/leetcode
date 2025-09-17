// 343. 整数拆分
// https://leetcode.cn/problems/integer-break

package integer_break

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		mid := i / 2
		v := i - 1
		for j := 2; j <= mid; j++ {
			v = max(v, max(j, dp[j])*max(i-j, dp[i-j]))
		}
		dp[i] = v
	}
	return dp[n]
}

func IntegerBreak(n int) int {
	return integerBreak(n)
}
