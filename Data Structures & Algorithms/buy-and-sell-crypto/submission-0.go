func maxProfit(prices []int) int {
	low, profit := prices[0],0
	for i:=0; i<len(prices); i++{
		low = min(low, prices[i])
		profit = max(profit, prices[i]-low)
	}
	return profit
}
