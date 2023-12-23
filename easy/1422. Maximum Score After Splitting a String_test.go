package easy

import (
	"math"
	"testing"

	"leetcode/uti"
)

// Time: O(n)
// Space: O(1)
func maxScore(s string) int {
	n := len(s)
	zeros, ones := 0, 0
	best := math.MinInt

	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones++
		}
		best = uti.Max(best, zeros-ones)
	}

	if s[n-1] == '1' {
		ones++
	}

	return best + ones
}

func TestMaxScore(t *testing.T) {
	for _, test := range []struct {
		s      string
		output int
	}{
		{s: "011101", output: 5},
		{s: "1111", output: 3},
	} {
		res := maxScore(test.s)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
