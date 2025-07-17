// 3176. 求出最长好子序列 I
// https://leetcode.cn/problems/find-the-maximum-length-of-a-good-subsequence-i/

package find_the_maximum_length_of_a_good_subsequence_i

func maximumLength(nums []int, k int) int {
	ans := 0
	left := 0
	sum := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		if right-1 >= left && nums[right-1] != nums[right] {
			sum++
		}
		for sum > k {
			if nums[left] != nums[left+1] {
				sum--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func MaximumLength(nums []int, k int) int {
	return maximumLength(nums, k)
}
