// 2110. 股票平滑下跌阶段的数目
// https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock

package number_of_smooth_descent_periods_of_a_stock

func getDescentPeriods(prices []int) int64 {
	var ans int64
	start := 0
	end := start + 1
	n := len(prices)
	for start < n {
		for end = start + 1; end < n && prices[end] == prices[end-1]-1; end++ {
		}
		m := int64(end - start)
		ans += m * (m + 1) / 2
		start = end
	}
	m := int64(end - start)
	ans += m * (m + 1) / 2
	return ans
}

func GetDescentPeriods(prices []int) int64 {
	return getDescentPeriods(prices)
}
