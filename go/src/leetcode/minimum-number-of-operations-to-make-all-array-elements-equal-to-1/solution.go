// 2654. 使数组所有元素变成 1 的最少操作次数
// https://leetcode.cn/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1

package minimum_number_of_operations_to_make_all_array_elements_equal_to_1

func minOperations(nums []int) int {
	n := len(nums)
	ds := make([]int, n)
	ones := 0
	for i, num := range nums {
		ds[i] = num
		if num == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			ds[j] = gcd(ds[j], ds[j+1])
			if ds[j] == 1 {
				return n + i
			}
		}
	}
	return -1
}

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func MinOperations(nums []int) int {
	return minOperations(nums)
}
