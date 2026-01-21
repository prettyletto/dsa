package arrays

func maxProfitBrute(prices []int) int {
	maxProfit := 0

	for i := range prices{
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i] 
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}


func maxProfit(prices[] int) int {
	l,r := 0, 1 
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = max(maxProfit, profit)
		}else{
			l = r
		}
		r ++
	}
	return maxProfit
}


