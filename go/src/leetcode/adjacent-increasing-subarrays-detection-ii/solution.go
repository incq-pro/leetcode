// 3350. 检测相邻递增子数组 II
// https://leetcode.cn/problems/adjacent-increasing-subarrays-detection-ii/

package adjacent_increasing_subarrays_detection_ii

func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	preCnt, curCnt := 0, 1
	ans := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			curCnt++
		} else {
			preCnt = curCnt
			curCnt = 1
		}
		ans = max(ans, min(preCnt, curCnt))
		ans = max(ans, curCnt/2)
	}
	return ans
}

func MaxIncreasingSubarrays(nums []int) int {
	return maxIncreasingSubarrays(nums)
}
