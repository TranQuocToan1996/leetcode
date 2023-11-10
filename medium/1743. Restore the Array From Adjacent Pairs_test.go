package medium

import (
	"testing"

	"leetcode/uti"
)

type inPair struct {
	first  *int
	second *int
}

// Time: O()
// Space: O()
func restoreArray(adjacentPairs [][]int) []int {
	var (
		m      = make(map[int]inPair, len(adjacentPairs)+1)
		result = make([]int, len(adjacentPairs)+1)
	)

	for _, pair := range adjacentPairs {
		markAdjacent(m, pair[0], pair[1])
		markAdjacent(m, pair[1], pair[0])
	}

	for key, pair := range m {
		if pair.second == nil {
			result[0] = key
			break
		}
	}

	for i := 1; i < len(result); i++ {
		cur := result[i-1]
		pair := m[cur]
		if pair.second == nil {
			result[i] = *pair.first
			continue
		}
		priv := result[i-2]
		if *pair.second == priv {
			result[i] = *pair.first
		} else {
			result[i] = *pair.second
		}
	}

	return result
}

func markAdjacent(m map[int]inPair, curVal, pairVal int) {
	if _, ok := m[curVal]; ok {
		m[curVal] = inPair{first: m[curVal].first, second: &pairVal}
	} else {
		m[curVal] = inPair{first: &pairVal}
	}
}

func TestRestoreArray(t *testing.T) {
	for _, test := range []struct {
		input  [][]int
		output []int
	}{
		{input: [][]int{
			{2, 1},
			{3, 4},
			{3, 2},
		}, output: []int{1, 2, 3, 4}},
		{input: [][]int{
			{4, -2},
			{1, 4},
			{-3, 1},
		}, output: []int{-2, 4, 1, -3}},
		{input: [][]int{
			{100000, -100000},
		}, output: []int{100000, -100000}},
		{input: [][]int{
			{100000, -100000},
		}, output: []int{-100000, 100000}},
		{input: [][]int{
			{4, -2},
			{1, 4},
			{-3, 1},
		}, output: []int{-2, 4, 1, -3}},
	} {
		res := restoreArray(test.input)
		if !uti.UnorderMatch[int](res, test.output) {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
