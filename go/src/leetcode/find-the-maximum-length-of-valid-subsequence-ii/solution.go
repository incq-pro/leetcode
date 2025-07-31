// 3202. 找出有效子序列的最大长度 II
// https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-ii/

package find_the_maximum_length_of_valid_subsequence_ii

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	res := 0
	for _, num := range nums {
		num %= k
		for prev := 0; prev < k; prev++ {
			dp[prev][num] = dp[num][prev] + 1
			res = max(res, dp[prev][num])
		}
	}
	return res
}

func MaximumLength(nums []int, k int) int {
	return maximumLength(nums, k)
}
