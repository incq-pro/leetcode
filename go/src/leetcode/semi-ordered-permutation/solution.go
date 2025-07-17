// 2717. 半有序排列
// https://leetcode.cn/problems/semi-ordered-permutation/

package semi_ordered_permutation

func semiOrderedPermutation(nums []int) int {
	a, b := -1, -1
	n := len(nums)
	for i, v := range nums {
		if v == 1 {
			a = i
		}
		if v == n {
			b = i
		}
	}
	r := a - 0 + (n - 1 - b)
	if a > b {
		r -= 1
	}
	return r
}

func SemiOrderedPermutation(nums []int) int {
	return semiOrderedPermutation(nums)
}
