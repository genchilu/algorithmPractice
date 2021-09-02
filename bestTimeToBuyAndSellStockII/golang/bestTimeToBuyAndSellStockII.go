package bestTimeToBuyAndSellStockII

func maxProfit(prices []int) int {

	val := prices[0]

	profile := 0
	for i := 1; i < len(prices); i++ {
		//fmt.Printf("%d, %d, %d\n", i, profile, val)
		if prices[i] < prices[i-1] {
			profile += prices[i-1] - val
			val = prices[i]
		} else if i == len(prices)-1 {
			profile += prices[i] - val
		}
	}
	return profile
}
