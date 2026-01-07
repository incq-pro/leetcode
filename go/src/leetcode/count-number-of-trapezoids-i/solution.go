// 3623. 统计梯形的数目 I
// https://leetcode.cn/problems/count-number-of-trapezoids-i

package count_number_of_trapezoids_i

func countTrapezoids(points [][]int) int {
	M := int(1e9 + 7)
	ans := 0
	cnt := make(map[int]int, len(points))
	for _, p := range points {
		cnt[p[1]]++
	}
	cnts := make([]int, 0, len(cnt))
	for _, v := range cnt {
		if v > 1 {
			cnts = append(cnts, v*(v-1)/2%M)
		}
	}
	n := len(cnts)
	for i := 1; i < n; i++ {
		cnts[i] = cnts[i-1] + cnts[i]
	}

	for i := n - 1; i > 0; i-- {
		ans += (cnts[i-1] * (cnts[i] - cnts[i-1])) % M
	}
	return ans % M
}

func CountTrapezoids(points [][]int) int {
	return countTrapezoids(points)
}
