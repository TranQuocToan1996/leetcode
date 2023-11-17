package medium

import (
	"math"
	"testing"

	"leetcode/uti"
)

// Time: O(n*2^n)
// Space: O(n*2^n)
// func subsets(nums []int) [][]int {
// 	n := len(nums)
// 	queue := make([][]int, 0, int(math.Pow(2, float64(n))))
// 	queue = append(queue, []int{})

// 	for i := 0; i < n; i++ {
// 		length := len(queue)
// 		for j := 0; j < length; j++ {
// 			temp := uti.DeepCopySlice(queue[j])
// 			temp = append(temp, nums[i])
// 			queue = append(queue, temp)
// 		}
// 	}

// 	return queue
// }

func subsets(nums []int) [][]int {
	n := len(nums)
	results := make([][]int, 0, int(math.Pow(2, float64(n))))
	subsetLength := 0
	var backtrack func(int, []int)
	backtrack = func(first int, cur []int) {
		if len(cur) == subsetLength {
			results = append(results, uti.DeepCopySlice(cur))
			return
		}
		for i := first; i < n; i++ {
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	for ; subsetLength < n+1; subsetLength++ {
		backtrack(0, []int{})
	}
	return results
}

// Time: O(2^n).
// Space: O(2^n)
/* func subsets(nums []int) [][]int {
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
} */

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
		// O(n^2). But n in the tests quite small.
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
