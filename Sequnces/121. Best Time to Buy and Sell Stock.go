package Sequnces

func maxProfit(prices []int) int {
	maxi := 0
	ans := 0
	for index := 1; index < len(prices); index++ {
		diff := prices[index] - prices[index-1]
		maxi = max(maxi+diff, 0)
		ans = max(ans, maxi)
	}
	return ans
}
