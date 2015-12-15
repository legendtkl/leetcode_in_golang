package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var ret int = 0
	var minv int = 1<<31 - 1
	for i := 0; i < len(prices); i++ {
		if minv > prices[i] {
			minv = prices[i]
		} else if ret < prices[i]-minv {
			ret = prices[i] - minv
		}
	}
	return ret
}

func main() {
	prices := []int{1, 9, 2, 11}
	fmt.Println(maxProfit(prices))
}
