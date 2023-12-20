package easy

import (
	"math"
	"testing"
)

// Time: O(n)
// Space: O(1)
func buyChoco(prices []int, money int) int {
	first, second := math.MaxInt, math.MaxInt
	for i := range prices {
		if first > prices[i] {
			second = first
			first = prices[i]
		} else if second > prices[i] {
			second = prices[i]
		}
	}
	if first+second > money {
		return money
	}
	return money - first - second
}

func TestBuyChoco(t *testing.T) {
	for _, test := range []struct {
		prices []int
		money  int
		output int
	}{
		{prices: []int{1, 2, 2}, money: 3, output: 0},
		{prices: []int{3, 2, 3}, money: 3, output: 3},
		{prices: []int{98, 54, 6, 34, 66, 63, 52, 39}, money: 62, output: 22},
		{prices: []int{69, 91, 78, 19, 40, 13}, money: 94, output: 62},
	} {
		res := buyChoco(test.prices, test.money)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
