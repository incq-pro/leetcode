// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/

package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	longest := 1
	for i := 1; i < n; i++ {
		li := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				li = max(li, dp[j])
			}
		}
		dp[i] = li + 1
		longest = max(longest, dp[i])
	}
	return longest
}

func LengthOfLIS(nums []int) int {
	return lengthOfLIS(nums)
}
