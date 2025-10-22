// 3349. 检测相邻递增子数组 I
// https://leetcode.cn/problems/adjacent-increasing-subarrays-detection-i

package adjacent_increasing_subarrays_detection_i

func hasIncreasingSubarrays(nums []int, k int) bool {
	preLen, curLen := 0, 0
	n := len(nums)
	for i := 0; i < n; {
		preLen = curLen
		curLen = 1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[j-1] {
				curLen++
			} else {
				break
			}
		}
		i += curLen
		if curLen >= 2*k || (curLen >= k && preLen >= k) {
			return true
		}
	}
	return false
}

func hasIncreasingSubarraysV2(nums []int, k int) bool {
	n := len(nums)
	preCnt, curCnt := 0, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			curCnt++
		} else {
			preCnt = curCnt
			curCnt = 1
		}
		if curCnt >= 2*k || (curCnt >= k && preCnt >= k) {
			return true
		}
	}
	return false
}

func HasIncreasingSubarrays(nums []int, k int) bool {
	return hasIncreasingSubarraysV2(nums, k)
}
