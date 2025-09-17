// 376. 摆动序列
// https://leetcode.cn/problems/wiggle-subsequence

package wiggle_subsequence

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	count := 1
	delta := 0
	i := 1
	for ; i < n; i++ {
		delta = nums[i] - nums[0]
		if delta != 0 {
			count++
			i++
			break
		}
	}
	for ; i < n; i++ {
		delta2 := nums[i] - nums[i-1]
		if delta2 == 0 {
			continue
		} else if delta2 > 0 {
			if delta < 0 {
				count++
				delta = delta2
			}
		} else {
			if delta > 0 {
				count++
				delta = delta2
			}
		}
	}

	return count
}

func WiggleMaxLength(nums []int) int {
	return wiggleMaxLength(nums)
}
