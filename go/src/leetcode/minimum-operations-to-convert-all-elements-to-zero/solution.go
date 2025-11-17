// 3542. 将所有元素变为 0 的最少操作次数
// https://leetcode.cn/problems/minimum-operations-to-convert-all-elements-to-zero

package minimum_operations_to_convert_all_elements_to_zero

func minOperationsLeetCode(nums []int) int {
	s := []int{}
	res := 0
	for _, a := range nums {
		for len(s) > 0 && s[len(s)-1] > a {
			s = s[:len(s)-1]
		}
		if a == 0 {
			continue
		}
		if len(s) == 0 || s[len(s)-1] < a {
			res++
			s = append(s, a)
		}
	}
	return res
}

func MinOperations(nums []int) int {
	return minOperationsLeetCode(nums)
}
