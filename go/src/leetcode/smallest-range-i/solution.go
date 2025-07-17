// 908. 最小差值 I
// https://leetcode.cn/problems/smallest-range-i/

package smallest_range_i

func smallestRangeI(nums []int, k int) int {
	// 先找到最大值和最小值，然后再计算分数
	n := len(nums)
	minV, maxV := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < minV {
			minV = nums[i]
		} else if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	delta := maxV - minV - 2*k
	if delta <= 0 {
		return 0
	}
	return delta
}

func SmallestRangeI(nums []int, k int) int {
	return smallestRangeI(nums, k)
}
