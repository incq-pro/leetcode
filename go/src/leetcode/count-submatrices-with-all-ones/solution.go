// 1504. 统计全 1 子矩形
// https://leetcode.cn/problems/count-submatrices-with-all-ones

package count_submatrices_with_all_ones

func numSubmat(mat [][]int) int {
	ans := 0
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			if j > 0 {
				mat[i][j] = mat[i][j-1] + 1
			}
			x := mat[i][j]
			for k := i; k >= 0; k-- {
				x = min(x, mat[k][j])
				if x == 0 {
					break
				}
				ans += x
			}
		}
	}
	return ans
}

func NumSubMat(mat [][]int) int {
	return numSubmat(mat)
}
