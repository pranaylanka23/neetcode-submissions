func maxProfit(prices []int) int {
	profit := 0
	low,high := prices[0], prices[0]
	for i:=1; i<len(prices);i++{
		if prices[i]<high{
			profit += high-low
			low = prices[i]
		}
		high = prices[i]
	}
	return profit+high-low
}
