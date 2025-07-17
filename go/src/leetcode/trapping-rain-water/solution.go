// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/

package trapping_rain_water

func trap(height []int) int {
	n := len(height)
	rightMax := make([]int, n)
	m := height[n-1]
	for i := n - 1; i >= 0; i-- {
		if height[i] > m {
			m = height[i]
		}
		rightMax[i] = m
	}

	m = height[0]
	s := 0
	for i, h := range height {
		if h > m {
			m = h
		}
		s += min(m, rightMax[i]) - h
	}
	return s
}

func Trap(height []int) int {
	return trap(height)
}
