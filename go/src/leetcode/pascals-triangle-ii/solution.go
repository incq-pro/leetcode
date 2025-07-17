// 119. 杨辉三角 II
// https://leetcode.cn/problems/pascals-triangle-ii

package pascals_triangle_ii

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		row[i] = 1
		for j := i - 1; j > 0; j-- {
			row[j] += row[j-1]
		}
		row[0] = 1
	}
	return row
}

func GetRow(rowIndex int) []int {
	return getRow(rowIndex)
}
