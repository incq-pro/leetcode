// https://leetcode.cn/problems/partition-equal-subset-sum
// 416. 分割等和子集

package partition_equal_subset_sum

func canPartition(nums []int) bool {
	s1, s2 := nums[0], 0
	left, right := 0, len(nums)
	for {
		if s1 < s2 {
			if left >= right-1 {
				break
			}
			left += 1
			s1 += nums[left]
		} else {
			if right <= left-1 {
				break
			}
			right -= 1
			s2 += nums[right]
		}
	}
	return s1 == s2
}

func CanPartition(nums []int) bool {
	return canPartition(nums)
}
