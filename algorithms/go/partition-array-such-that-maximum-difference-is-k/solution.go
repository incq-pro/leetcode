// 2294. 划分数组使最大差为 K
// https://leetcode.cn/problems/partition-array-such-that-maximum-difference-is-k/

package partition_array_such_that_maximum_difference_is_k

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	c := 0
	for e, s := 0, 0; s < n && e < n; {
		t := nums[s] + k
		for e < n && nums[e] <= t {
			e++
		}
		c += 1
		s = e
	}
	return c
}

func PartitionArray(nums []int, k int) int {
	return partitionArray(nums, k)
}
