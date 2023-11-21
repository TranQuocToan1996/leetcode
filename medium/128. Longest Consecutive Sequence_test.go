package medium

import (
	"testing"

	"leetcode/uti"
)

// Time: O(n)
// Space: O(n)
// You must write an algorithm that runs in O(n) time.
func longestConsecutive(nums []int) int {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		seen[n] = struct{}{}
	}

	max := 0
	for n := range seen {
		if _, ok := seen[n-1]; !ok {
			seqLen := 1
			for {
				if _, ok = seen[n+seqLen]; ok {
					seqLen++
					continue
				}
				max = uti.Max(seqLen, max)
				break
			}
		}
	}

	return max
}

func TestLongestConsecutive(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, output: 7},
		{nums: []int{100, 4, 200, 1, 3, 2}, output: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, output: 9},
		{nums: []int{}, output: 0},
	} {
		res := longestConsecutive(test.nums)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
