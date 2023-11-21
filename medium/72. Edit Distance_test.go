package medium

import (
	"testing"

	"leetcode/uti"
)

// Time: O(n*m)
// Space: O(n*m)
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	cache := make([][]int, n1+1)
	for i := range cache {
		cache[i] = make([]int, n2+1)
	}

	for i := 1; i <= n2; i++ {
		cache[0][i] = i
	}
	for i := 1; i <= n1; i++ {
		cache[i][0] = i
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				cache[i][j] = cache[i-1][j-1]
				continue
			}
			cache[i][j] = uti.Min(
				uti.Min(
					cache[i-1][j],
					cache[i][j-1],
				), cache[i-1][j-1]) + 1
		}
	}

	return cache[n1][n2]
}

func TestMinDistance(t *testing.T) {
	for _, test := range []struct {
		start  string
		end    string
		output int
	}{
		{start: "horse", end: "ros", output: 3},
		{start: "intention", end: "execution", output: 5},
	} {
		res := minDistance(test.start, test.end)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}


