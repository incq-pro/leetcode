// 3195. 包含所有 1 的最小矩形面积 I
// https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/

package find_the_minimum_area_to_cover_all_ones_i

func minimumArea(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	top, bottom, left, right := m, -1, n, -1
	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i < top {
				top = i
			}
			if i > bottom {
				bottom = i
			}
			if j < left {
				left = j
			}
			if j > right {
				right = j
			}
		}
	}

	return (bottom - top + 1) * (right - left + 1)
}

func MinimumArea(grid [][]int) int {
	return minimumArea(grid)
}
