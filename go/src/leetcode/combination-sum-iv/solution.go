// 377. 组合总和 Ⅳ
// https://leetcode.cn/problems/combination-sum-iv/

package combination_sum_iv

import "sort"

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		cnt := 0
		for j := 0; j < len(nums) && nums[j] <= i; j++ {
			cnt += dp[i-nums[j]]
		}
		dp[i] = cnt
	}
	return dp[target]
}

func CombinationSum4(nums []int, target int) int {
	return combinationSum4(nums, target)
}
