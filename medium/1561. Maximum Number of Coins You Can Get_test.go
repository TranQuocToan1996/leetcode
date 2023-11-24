package medium

import (
	"sort"
	"testing"
)

// Time: O(nlogn)
// Space: O(logn)
func maxCoins(piles []int) int {
	var (
		max  = 0
		step = 3 // Bob, you, Alice
	)
	sort.IntSlice(piles).Sort()
	piles = piles[len(piles)/step:] // Remove Bob
	step--                          // You and Alice
	for i := 0; i < len(piles); i += step {
		max += piles[i]
	}
	return max
}

func TestMaxCoins(t *testing.T) {
	for _, test := range []struct {
		piles  []int
		output int
	}{
		{piles: []int{9, 8, 7, 6, 5, 1, 2, 3, 4}, output: 18},
		{piles: []int{2, 4, 1, 2, 7, 8}, output: 9},
		{piles: []int{2, 4, 5}, output: 4},
	} {
		res := maxCoins(test.piles)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
