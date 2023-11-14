package medium

import (
	"reflect"
	"testing"
)

// Time: O(m*n)
// Space: O(m*n)
func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		cache[m-1][i] = 1
	}

	for i := 0; i < m; i++ {
		cache[i][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			cache[i][j] = cache[i+1][j] + cache[i][j+1]
		}
	}

	return cache[0][0]
}

func TestUniquePaths(t *testing.T) {
	for _, test := range []struct {
		m      int
		n      int
		output int
	}{
		{
			m:      3,
			n:      7,
			output: 28,
		},
		{
			m:      3,
			n:      2,
			output: 3,
		},
	} {
		res := uniquePaths(test.m, test.n)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test)
		}
	}
}
