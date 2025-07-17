// 2708. 一个小组的最大实力值
// https://leetcode.cn/problems/maximum-strength-of-a-group/

package maximum_strength_of_a_group

import "slices"

func maxStrength(nums []int) int64 {
	slices.Sort(nums)
	r := make([]int, 0, len(nums))
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			r = append(r, i)
		} else if nums[i] < 0 {
			if i+1 < len(nums) && nums[i+1] < 0 {
				r = append(r, i, i+1)
				i++
			}
		} else if nums[i] == 0 {
			zeroIndex = i
		}
	}
	if len(r) == 0 {
		if zeroIndex >= 0 {
			r = append(r, zeroIndex)
		} else {
			r = append(r, 0)
		}
	}
	var result int64 = 1
	for _, i := range r {
		result *= int64(nums[i])
	}
	return result
}

func MaxStrength(nums []int) int64 {
	return maxStrength(nums)
}
