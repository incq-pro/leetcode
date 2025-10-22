// 1039. 多边形三角剖分的最低得分
// https://leetcode.cn/problems/minimum-score-triangulation-of-polygon

package minimum_score_triangulation_of_polygon

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 2; l < n; l++ {
		for i := 0; i < n-1; i++ {
			j := (i + 2) % n
			v := 0
			for k := 1; k < l; k++ {
				for m := 1; m <= l-k; m++ {
					ik := (i + k) % n
					im := (i + k + m) % n
					v = max(v, values[i]*values[ik]*values[im]+dp[i][ik]+dp[ik][im]+dp[im][i])
				}
			}
			dp[i][j] = v
		}
	}
	return dp[0][n-1]
}

func MinScoreTriangulation(values []int) int {
	return minScoreTriangulation(values)
}
