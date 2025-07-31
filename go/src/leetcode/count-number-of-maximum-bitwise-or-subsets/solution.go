// 2044. 统计按位或能得到最大值的子集数目
// https://leetcode.cn/problems/count-number-of-maximum-bitwise-or-subsets/

package count_number_of_maximum_bitwise_or_subsets

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}

	ans := 0
	for i := 1; i < 1<<n; i++ {
		or := 0
		for j, num := range nums {
			if (i>>j)&1 == 1 {
				or |= num
			}
		}
		if or == maxOr {
			ans++
		}
	}
	return ans
}

func CountMaxOrSubsets(nums []int) int {
	return countMaxOrSubsets(nums)
}
