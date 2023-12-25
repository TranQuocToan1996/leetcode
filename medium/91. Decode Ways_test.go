package medium

import (
	"testing"
)

// time: O(n)
// space: O(n)
// func numDecodings(s string) int {
// 	n := len(s)
// 	dp := map[int]int{n: 1}

// 	for i := n - 1; i >= 0; i-- {
// 		if s[i] == '0' {
// 			dp[i] = 0
// 		} else {
// 			dp[i] = dp[i+1]
// 		}

// 		if canNumDecoding(s, i) {
// 			dp[i] += dp[i+2]
// 		}
// 	}

// 	return dp[0]
// }

// func canNumDecoding(s string, i int) bool {
// 	return i+1 < len(s) && s[i:i+2] >= "10" && s[i:i+2] <= "26"
// }

// time: O(n)
// space: O(1)
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	prev, cur := 1, 1
	for i := 1; i < n; i++ {
		next := 0
		if s[i] != '0' {
			next += cur
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			next += prev
		}
		prev, cur = cur, next
	}

	return cur
}

func TestNumDecodings(t *testing.T) {
	for _, test := range []struct {
		s      string
		output int
	}{
		{s: "12", output: 2},
		{s: "226", output: 3},
		{s: "06", output: 0},
		{s: "27", output: 1},
	} {
		res := numDecodings(test.s)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
