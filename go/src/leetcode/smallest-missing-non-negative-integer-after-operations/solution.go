// 2598. 执行操作后的最大 MEX
// https://leetcode.cn/problems/smallest-missing-non-negative-integer-after-operations

package smallest_missing_non_negative_integer_after_operations

func findSmallestInteger(nums []int, value int) int {
	reminder := make([]int, value)
	for _, num := range nums {
		r := num % value
		if r < 0 {
			r = r + value
		}
		reminder[r] = reminder[r] + 1
	}
	m := 0
	for i := 0; i < value; i++ {
		if reminder[i] < reminder[m] {
			m = i
		}
	}
	return reminder[m]*value + m
}

func FindSmallestInteger(nums []int, value int) int {
	return findSmallestInteger(nums, value)
}
