// 2348. 全 0 子数组的数目
// https://leetcode.cn/problems/number-of-zero-filled-subarrays

package number_of_zero_filled_subarrays

func zeroFilledSubarray(nums []int) int64 {
	var ret int64
	var c int64
	for _, v := range nums {
		if v == 0 {
			c++
			ret += c
		} else {
			c = 0
		}
	}
	return ret
}

func ZeroFilledSubarray(nums []int) int64 {
	return zeroFilledSubarray(nums)
}
