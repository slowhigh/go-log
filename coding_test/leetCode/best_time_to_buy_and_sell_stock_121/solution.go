package best_time_to_buy_and_sell_stock_121

func maxProfit(prices []int) int {
	min, max := prices[0], prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			max = prices[i]
		} else if max < prices[i] {
			max = prices[i]
		}

		if profit < max - min {
			profit = max - min
		}
	}

	return profit
}
