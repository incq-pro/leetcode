// https://leetcode.cn/problems/check-if-all-1s-are-at-least-length-k-places-away/
// 1437. 是否所有 1 都至少相隔 k 个元素

package check_if_all_1s_are_at_least_length_k_places_away

func kLengthApart(nums []int, k int) bool {
	last := -1
	for i, num := range nums {
		if num == 1 {
			last = i
			break
		}
	}
	if last < 0 {
		return true
	}
	for i := last + 1; i < len(nums); i++ {
		if nums[i] == 1 {
			if i-last-1 < k {
				return false
			}
			last = i
		}
	}
	return true
}

func KLengthApart(nums []int, k int) bool {
	return kLengthApart(nums, k)
}
