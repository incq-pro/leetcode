// 2073. 买票需要的时间
// https://leetcode.cn/problems/time-needed-to-buy-tickets

package time_needed_to_buy_tickets

func timeRequiredToBuy(tickets []int, k int) int {
	total := 0
	for i := 0; i <= k; i++ {
		total += min(tickets[i], tickets[k])
	}
	ktickets := tickets[k] - 1
	for i := k + 1; i < len(tickets); i++ {
		total += min(tickets[i], ktickets)
	}
	return total
}

func TimeRequiredToBuy(tickets []int, k int) int {
	return timeRequiredToBuy(tickets, k)
}
