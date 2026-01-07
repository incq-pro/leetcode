// 3381. 长度可被 K 整除的子数组的最大元素和
// https://leetcode.cn/problems/maximum-subarray-sum-with-length-divisible-by-k

package maximum_subarray_sum_with_length_divisible_by_k

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := int64(0)
	maxSum := int64(math.MinInt64)
	kSum := make([]int64, k)
	for i := 0; i < k-1; i++ {
		kSum[i] = math.MaxInt64
	}
	kSum[k-1] = 0
	for i := 0; i < n; i++ {
		prefixSum += int64(nums[i])
		if i+1 >= k {
			maxSum = max(maxSum, prefixSum-kSum[i%k])
		}
		kSum[i%k] = min(kSum[i%k], prefixSum)
	}
	return maxSum
}

func MaxSubarraySum(nums []int, k int) int64 {
	return maxSubarraySum(nums, k)
}
