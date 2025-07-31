// 2419. 按位与最大的最长子数组
// https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and

package longest_subarray_with_maximum_bitwise_and

func longestSubarray0(nums []int) int {
	maxV := nums[0]
	maxLen := 1
	vv := nums[0]
	ll := 0
	for _, v := range nums {
		if v == vv {
			ll++
		} else {
			if vv > maxV {
				maxV = vv
				maxLen = ll
			} else if vv == maxV && ll > maxLen {
				maxLen = ll
			}
			vv = v
			ll = 1
		}
	}
	if vv > maxV {
		maxV = vv
		maxLen = ll
	} else if vv == maxV && ll > maxLen {
		maxLen = ll
	}
	return maxLen
}

func longestSubarray(nums []int) int {
	maxValue := nums[0]
	ans, cnt := 1, 1
	n := len(nums)
	for i := 1; i < n; i++ {
		v := nums[i]
		if v > maxValue {
			maxValue = v
			ans, cnt = 1, 1
		} else if v == maxValue {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 0
		}
	}
	return ans
}

func LongestSubarray(nums []int) int {
	return longestSubarray(nums)
}
