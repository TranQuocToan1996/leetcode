package medium

import (
	"reflect"
	"strings"
	"testing"
)

// Time: O(n)
// Space: O(1) 26 chars only in seen map
func partitionLabels(s string) []int {
	const notfound = -1
	var (
		res    []int
		seen   = map[byte]bool{}
		greedy func(s string)
	)
	greedy = func(s string) {
		if len(s) == 0 {
			return
		}
		left, right := string(s[0]), s[1:]
		for i := 0; i < len(left); i++ {
			if seen[left[i]] {
				continue
			}
			seen[left[i]] = true
			lastIndex := strings.LastIndex(right, string(left[i]))
			if lastIndex == notfound {
				continue
			}
			left, right = left+right[:lastIndex+1], right[lastIndex+1:]
		}
		res = append(res, len(left))
		greedy(right)
	}
	greedy(s)
	return res
}

func TestPartitionLabels(t *testing.T) {
	for _, test := range []struct {
		s      string
		output []int
	}{
		{s: "ababcbacadefegdehijhklij", output: []int{9, 7, 8}},
		{s: "eccbbbbdec", output: []int{10}},
		{s: "caedbdedda", output: []int{1, 9}},
	} {
		res := partitionLabels(test.s)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
