// 1975. 最大方阵和
// https://leetcode.cn/problems/maximum-matrix-sum

package maximum_matrix_sum

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	m := len(matrix)
	n := len(matrix[0])
	// 负数个数
	nc := 0
	// 需要扣除的最小值
	var x int64 = math.MaxInt64
	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := int64(matrix[i][j])
			if c >= 0 {
				ans += c
				x = min(x, c)
			} else {
				ans += -c
				x = min(x, -c)
				nc++
			}
		}
	}
	if nc%2 != 0 {
		ans -= x * 2
	}
	return ans
}

func MaxMatrixSum(matrix [][]int) int64 {
	return maxMatrixSum(matrix)
}
