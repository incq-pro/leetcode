// 3239. 最少翻转次数使二进制矩阵回文 I
// https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/

package minimum_number_of_flips_to_make_binary_grid_palindromic_i

func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	c1, c2 := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n-1-j; j++ {
			if grid[i][j] != grid[i][n-1-j] {
				c1++
			}
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m-1-i; i++ {
			if grid[i][j] != grid[m-1-i][j] {
				c2++
			}
		}
	}
	return min(c1, c2)
}

func MinFlips(grid [][]int) int {
	return minFlips(grid)
}
