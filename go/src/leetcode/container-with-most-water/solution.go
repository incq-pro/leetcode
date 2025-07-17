// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/

package container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	m := 0
	for left < right {
		if height[left] > height[right] {
			m = max(m, (right-left)*height[right])
			right--
		} else {
			m = max(m, (right-left)*height[left])
			left++
		}
	}
	return m
}

func MaxArea(height []int) int {
	return maxArea(height)
}
