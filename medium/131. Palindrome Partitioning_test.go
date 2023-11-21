package medium

import (
	"reflect"
	"testing"

	"leetcode/uti"
)

// Time: O(2^n * n)
// Space: O(n)
func partition(s string) [][]string {
	n := len(s)
	var res [][]string
	var part []string
	var backtrack func(int)
	backtrack = func(start int) {
		if start == n {
			res = append(res, uti.DeepCopySlice(part))
			return
		}

		for end := start; end < n; end++ {
			if uti.IsPalindrome(s[start : end+1]) {
				part = append(part, s[start:end+1])
				backtrack(end + 1)
				part = part[:len(part)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func TestPartition(t *testing.T) {
	for _, test := range []struct {
		start  string
		output [][]string
	}{
		{start: "aab", output: [][]string{
			{"a", "a", "b"},
			{"aa", "b"},
		}},
		{start: "a", output: [][]string{
			{"a"},
		}},
	} {
		res := partition(test.start)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
