package main

import (
	"fmt"
)

func main() {

	tests := [][]int{
		{8, 4, 6, 2, 3},
		{10, 1, 1, 6},
		{1, 2, 3, 4, 5},
	}
	for _, test := range tests {
		fmt.Println(finalPrices(test))
	}
}

// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
/*
time complexity: O(n*n) 2 for loop, 3 ms
Space complexity: O(1), not too much memory generation in this code, 3 MB
*/
func finalPrices(prices []int) []int {
	length := len(prices)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if prices[i] >= prices[j] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}
