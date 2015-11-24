// solution: DP
// status: AC
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	} else {
		dp := make([]int, n)
		dp[0] = 0
		dp[1] = max(prices[1]-prices[0], 0)
		for i := 2; i < n; i++ {
			for j := i - 1; j > 1; j-- {
				dp[i] = max(dp[i], max(prices[i]-prices[j], 0)+dp[j-2])
			}
			dp[i] = max(dp[i], max(prices[i]-prices[1], prices[i]-prices[0]))
			dp[i] = max(dp[i-1], dp[i])
		}
		return dp[n-1]
	}
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
