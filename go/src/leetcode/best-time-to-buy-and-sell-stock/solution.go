// 121. 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	lowest := prices[0]
	result := 0
	for _, v := range prices {
		result = max(result, v-lowest)
		lowest = min(lowest, v)
	}
	return result
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}
