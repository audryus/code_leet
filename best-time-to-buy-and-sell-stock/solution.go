package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		profit := price - int(minPrice)

		if profit > maxProfit {
			maxProfit = profit
		}

	}

	return maxProfit
}
