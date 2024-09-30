package easy

import "testing"

func TestMaxProfit(t *testing.T) {
	for _, test := range []struct {
		prices []int
		output int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, output: 5},
		{prices: []int{7, 6, 4, 3, 1}, output: 0},
	} {
		if maxProfit(test.prices) != test.output {
			t.Errorf("Expect %v but got %v at %v", test.output, maxProfit(test.prices), test.prices)
		}
	}
}

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for _, price := range prices {
		if buy > price {
			buy = price
		} else if profit < price-buy {
			profit = price - buy
		}
	}
	return profit
}
