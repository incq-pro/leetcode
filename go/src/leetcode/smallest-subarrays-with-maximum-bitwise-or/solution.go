// 2411. 按位或最大的最小子数组长度
// https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/

package smallest_subarrays_with_maximum_bitwise_or

func smallestSubarrays(nums []int) []int {
	index := make([]int, 32)
	ans := make([]int, len(nums))
	maxOr := 0
	for i := len(nums) - 1; i >= 0; i-- {
		maxOr |= nums[i]
		maxIndex := i
		for j := len(index) - 1; j >= 0; j-- {
			if (nums[i]>>j)&1 == 1 {
				index[j] = i
			} else if ((maxOr>>j)&1) == 1 && maxIndex < index[j] {
				maxIndex = index[j]
			}
		}
		ans[i] = maxIndex - i + 1
	}
	return ans
}

func SmallestSubarrays(nums []int) []int {
	return smallestSubarrays(nums)
}
