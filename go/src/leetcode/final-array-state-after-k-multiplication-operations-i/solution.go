// 3264. K 次乘运算后的最终数组 I
// https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-i

package final_array_state_after_k_multiplication_operations_i

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n := len(nums)
	for i := 0; i < k; i++ {
		k := 0
		for j := 1; j < n; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[k] *= multiplier
	}
	return nums
}
func GetFinalState(nums []int, k int, multiplier int) []int {
	return getFinalState(nums, k, multiplier)
}
