// 2270. 分割数组的方案数
// https://leetcode.cn/problems/number-of-ways-to-split-array/

package number_of_ways_to_split_array

func waysToSplitArray(nums []int) int {
	var rightSum int64
	for _, v := range nums {
		rightSum += int64(v)
	}
	count := 0
	var leftSum int64 = int64(nums[len(nums)-1])
	rightSum -= leftSum
	for i := len(nums) - 2; i >= 0; i-- {
		if rightSum >= leftSum {
			count++
		}
		rightSum -= int64(nums[i])
		leftSum += int64(nums[i])
	}
	return count
}

func WaysToSplitArray(nums []int) int {
	return waysToSplitArray(nums)
}
