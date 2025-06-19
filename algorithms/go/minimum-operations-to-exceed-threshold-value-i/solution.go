// 3065. 超过阈值的最少操作数 I
// https://leetcode.cn/problems/minimum-operations-to-exceed-threshold-value-i/

package minimum_operations_to_exceed_threshold_value_i

func minOperations(nums []int, k int) int {
	count := 0
	for _, v := range nums {
		if v < k {
			count++
		}
	}
	return count
}

func MinOperations(nums []int, k int) int {
	return minOperations(nums, k)
}
