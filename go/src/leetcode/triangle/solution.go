// 120. 三角形最小路径和
// https://leetcode.cn/problems/triangle

package triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	last := make([]int, n)
	last[0] = triangle[0][0]
	cur := make([]int, n)

	for row := 1; row < n; row++ {
		cur[0] = triangle[row][0] + last[0]
		cur[row] = triangle[row][row] + last[row-1]
		for j := 1; j < row; j++ {
			cur[j] = triangle[row][j] + min(last[j-1], last[j])
		}
		cur, last = last, cur
	}
	minV := last[0]
	for i := 1; i < n; i++ {
		if last[i] < minV {
			minV = last[i]
		}
	}
	return minV
}

func MinimumTotal(triangle [][]int) int {
	return minimumTotal(triangle)
}
