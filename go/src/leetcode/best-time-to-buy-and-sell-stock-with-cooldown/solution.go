// 309. 买卖股票的最佳时机含冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/

package best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	f0, f1, f2 := -prices[0], 0, 0
	n := len(prices)
	for i := 1; i < n; i++ {
		newF0 := max(f0, f2-prices[i])
		newF1 := f0 + prices[i]
		newF2 := max(f2, f1)
		f0, f1, f2 = newF0, newF1, newF2
	}
	return max(f1, f2)
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}
