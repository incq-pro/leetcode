// 3095. 或值至少 K 的最短子数组 I
// https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-i/

package shortest_subarray_with_or_at_least_k_i

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	result := -1
	left, right := 0, 0
	for right < n {
		b := 0
		for ; right < n; right++ {
			b = allOr(nums, left, right)
			if b >= k {
				break
			}
		}
		if b >= k {
			for ; left <= right; left++ {
				b = allOr(nums, left, right)
				if b < k {
					break
				}
			}
			if result == -1 {
				result = right - left + 2
			} else {
				result = min(result, right-left+2)
			}
			right++
		}
	}
	return result
}

func allOr(nums []int, left int, right int) int {
	b := 0
	for i := left; i <= right; i++ {
		b |= nums[i]
	}
	return b
}

func MinimumSubarrayLength(nums []int, k int) int {
	return minimumSubarrayLength(nums, k)
}
