package medium

import (
	"math"
	"testing"

	"leetcode/uti"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	results := make([][]int, 0, int(math.Pow(2, float64(n))))
	results = append(results, []int{})
	var generate func([]int)
	generate = func(s []int) {
		match := false
		for _, result := range results {
			if uti.UnorderMatch(result, s) {
				match = true
				break
			}
		}
		if !match {
			results = append(results, s)
		}

		if len(s) == n {
			return
		}
		for _, num := range nums {
			// TODO: Check duplicate nums
			for _, val := range s {
				if num == val {
					return
				}
			}
			newSlice := uti.DeepCopySlice(append(s, num))
			generate(newSlice)
		}
	}
	generate([]int{})
	return results
}

func TestSubsets(t *testing.T) {
	for _, test := range []struct {
		input  []int
		output [][]int
	}{
		{
			input: []int{1, 2, 3},
			output: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{2, 3},
				{3, 1},
				{1, 2, 3},
			},
		},
		{
			input: []int{0},
			output: [][]int{
				{},
				{0},
			},
		},
	} {
		match := 0
		output := subsets(test.input)
		// O(n^2). But n in the tests quite small, so accept this time.
		for _, subsetOutput := range test.output {
			for _, subset := range output {
				if uti.UnorderMatch(subset, subsetOutput) {
					match++
				}
			}
		}
		if match != len(test.output) {
			t.Errorf("expect %v but got %v, test [%v]", test.output, output, test.input)
		}
	}
}
