// 357. 统计各位数字都不同的数字个数
// https://leetcode.cn/problems/count-numbers-with-unique-digits

package count_numbers_with_unique_digits

func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	if n >= 1 {
		dp[1] = 10
	}
	r := 9
	for i := 2; i <= n; i++ {
		r *= (11 - i)
		dp[i] = dp[i-1] + r
	}
	return dp[n]
}

func CountNumbersWithUniqueDigits(n int) int {
	return countNumbersWithUniqueDigits(n)
}
