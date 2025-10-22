// 611. 有效三角形的个数
// https://leetcode.cn/problems/valid-triangle-number/

package valid_triangle_number

import "slices"

func triangleNumber(nums []int) int {
	slices.Sort(nums)
	c := 0
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] == 0 {
			continue
		}
		j := i + 1
		k := j + 1
		for ; j < n-1; j++ {
			k = max(k, j+1)
			for ; k < n && nums[k] < nums[i]+nums[j]; k++ {
			}
			c += k - 1 - j
		}
	}
	return c
}

func TriangleNumber(nums []int) int {
	return triangleNumber(nums)
}
