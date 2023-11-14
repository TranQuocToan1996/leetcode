package medium

import (
	"strings"
	"testing"
)

// Method 1
// Time: O(n)
// Space: O(1) only 26 unique chars for seen/unique map
/* func countPalindromicSubsequence(s string) int {
	var (
		seen   = make(map[byte]bool)
		result = 0
	)
	for first := 0; first < len(s)-2; first++ {
		if seen[s[first]] {
			continue
		}
		for end := len(s) - 1; end > first+1; end-- {
			if s[first] == s[end] {
				result += countUniqueChar(s[first+1 : end])
				break
			}
		}
		seen[s[first]] = true
	}

	return result
} */

func countUniqueChar(s string) int {
	unique := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		unique[s[i]] = true
	}
	return len(unique)
}

// Method 2
func countPalindromicSubsequence(s string) int {
	var (
		seen   = make(map[byte]bool)
		result = 0
	)
	for i := 0; i < len(s)-1; i++ {
		if seen[s[i]] {
			continue
		}
		first, end := strings.Index(s, string(s[i])), strings.LastIndex(s, string(s[i]))
		if first != end {
			result += countUniqueChar(s[first+1 : end])
		}
		seen[s[i]] = true
	}

	return result
}

// Method 3
/* func countPalindromicSubsequence(s string) int {
	const notFound = -1
	var (
		// Space O(1)
		seenFirst = [26]int{}
		seenEnd   = [26]int{}
		result    = 0
	)
	for i := 0; i < len(seenFirst); i++ {
		seenFirst[i] = notFound
		seenEnd[i] = notFound
	}

	// Time O(n)
	for i := 0; i < len(s); i++ {
		key := 'z' - s[i]
		first, end := strings.Index(s, string(s[i])), strings.LastIndex(s, string(s[i]))
		if first != end {
			seenFirst[key] = first
			seenEnd[key] = end
		}
	}

	for i := 0; i < len(seenFirst); i++ {
		first, end := seenFirst[i], seenEnd[i]
		if first != end {
			result += countUniqueChar(s[first+1 : end])
		}
	}

	return result
} */

func TestCountPalindromicSubsequence(t *testing.T) {
	for _, test := range []struct {
		input  string
		output int
	}{
		{input: "aabca", output: 3},
		{input: "adc", output: 0},
		{input: "bbcbaba", output: 4},
	} {
		res := countPalindromicSubsequence(test.input)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
