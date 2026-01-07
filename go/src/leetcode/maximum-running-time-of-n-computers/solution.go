// 2141. 同时运行 N 台电脑的最长时间
// https://leetcode.cn/problems/maximum-running-time-of-n-computers/

package maximum_running_time_of_n_computers

func maxRunTime(n int, batteries []int) int64 {
	var ans int64
	var s int64
	for _, b := range batteries {
		s += int64(b)
	}
	left, right := int64(0), s/int64(n)
	for left <= right {
		mid := left + (right-left)/2
		total := int64(0)
		for _, b := range batteries {
			total += min(int64(b), mid)
		}
		if total >= mid*int64(n) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func MaxRunTime(n int, batteries []int) int64 {
	return maxRunTime(n, batteries)
}
