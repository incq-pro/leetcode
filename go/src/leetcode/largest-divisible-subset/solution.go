// 368. 最大整除子集
// https://leetcode.cn/problems/largest-divisible-subset

package largest_divisible_subset

import "slices"

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	dp := make([][]int, len(nums))
	n := len(nums)
	mi := 0
	dp[0] = []int{1, -1}
	for i := 1; i < n; i++ {
		dp[i] = []int{1, -1}
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if dp[j][0]+1 > dp[i][0] {
					dp[i][0] = dp[j][0] + 1
					dp[i][1] = j
				}
			}
		}
		if dp[i][0] > dp[mi][0] {
			mi = i
		}
	}
	ret := make([]int, 0, dp[mi][0])
	ret = append(ret, nums[mi])
	for dp[mi][1] >= 0 {
		mi = dp[mi][1]
		ret = append(ret, nums[mi])
	}
	return ret
}

func LargestDivisibleSubset(nums []int) []int {
	return largestDivisibleSubset(nums)
}
