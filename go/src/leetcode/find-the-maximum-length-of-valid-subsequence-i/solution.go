// 3201. 找出有效子序列的最大长度 I
// https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-i/

package find_the_maximum_length_of_valid_subsequence_i

func maximumLength(nums []int) int {
	c0, c1, c2 := 0, 0, 0
	r := -1
	for _, num := range nums {
		rr := num % 2
		if rr == 0 {
			c0 += 1
		} else {
			c1 += 1
		}
		if rr != r {
			c2 += 1
			r = rr
		}
	}
	return max(max(c0, c1), c2)
}

func MaximumLength(nums []int) int {
	return maximumLength(nums)
}
