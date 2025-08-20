// 1277. 统计全为 1 的正方形子矩阵
// https://leetcode.cn/problems/count-square-submatrices-with-all-ones

package count_square_submatrices_with_all_ones

func countSquares(matrix [][]int) int {
	ans := 0
	m := len(matrix)
	n := len(matrix[0])

	for _, v := range matrix[0] {
		ans += v
	}
	for j := 1; j < m; j++ {
		ans += matrix[j][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := matrix[i][j]
			if v > 0 {
				v = min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1]) + 1
				ans += v
				matrix[i][j] = v
			}
		}
	}
	return ans
}

func CountSquares(matrix [][]int) int {
	return countSquares(matrix)
}
