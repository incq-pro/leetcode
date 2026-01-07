// 840. 矩阵中的幻方
// https://leetcode.cn/problems/magic-squares-in-grid

package magic_squares_in_grid

func numMagicSquaresInside(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}
	ans := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if isMagic(grid, i, j) {
				ans += 1
			}
		}
	}
	return ans
}

func isMagic(grid [][]int, top, left int) bool {
	h := make([]int, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[top+i][left+j] > 9 || grid[top+i][left+j] < 1 {
				return false
			} else if h[grid[top+i][left+j]] != 0 {
				return false
			}
			h[grid[top+i][left+j]] = 1
		}
	}
	// 每行
	for i := 0; i < 3; i++ {
		if (grid[top+i][left] + grid[top+i][left+1] + grid[top+i][left+2]) != 15 {
			return false
		}
	}
	// 每列
	for j := 0; j < 3; j++ {
		if (grid[top][left+j] + grid[top+1][left+j] + grid[top+2][left+j]) != 15 {
			return false
		}
	}
	// 斜对角
	if (grid[top][left] + grid[top+1][left+1] + grid[top+2][left+2]) != 15 {
		return false
	}
	if (grid[top][left+2] + grid[top+1][left+1] + grid[top+2][left]) != 15 {
		return false
	}
	return true
}

func NumMagicSquaresInside(grid [][]int) int {
	return numMagicSquaresInside(grid)
}
