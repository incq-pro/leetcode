// 122. 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	result := 0
	n := len(prices)
	for i := 0; i < n-1; i++ {
		delta := prices[i+1] - prices[i]
		if delta > 0 {
			result += delta
		}
	}
	return result
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}
