// 976. 三角形的最大周长
// https://leetcode.cn/problems/largest-perimeter-triangle

package largest_perimeter_triangle

import "slices"

func largestPerimeter(nums []int) int {
	slices.Sort(nums)
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func LargestPerimeter(nums []int) int {
	return largestPerimeter(nums)
}
