// 1018. 可被 5 整除的二进制前缀
// https://leetcode.cn/problems/binary-prefix-divisible-by-5

package binary_prefix_divisible_by_5

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	ans := make([]bool, n)
	s := 0
	for i := 0; i < n; i++ {
		s = (s*2 + nums[i]) % 10
		ans[i] = (s % 5) == 0
	}
	return ans
}

func PrefixesDivBy5(nums []int) []bool {
	return prefixesDivBy5(nums)
}
