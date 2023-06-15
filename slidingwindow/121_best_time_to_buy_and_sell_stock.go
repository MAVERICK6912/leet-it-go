package slidingwindow

func maxProfit(prices []int) int {
	if len(prices) > 1 {
		left, right := 0, 1
		maxProfit := 0
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		for right < len(prices) {
			if prices[left] < prices[right] {
				maxProfit = max(maxProfit, prices[right]-prices[left])
			} else {
				left = right
			}
			right += 1
		}
		return maxProfit
	}
	return 0
}
