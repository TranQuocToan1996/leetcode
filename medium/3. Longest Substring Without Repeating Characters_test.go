package medium

import (
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range []struct {
		input   string
		epected int
	}{
		{input: "abcabcbb", epected: 3},
		{input: "bbbbb", epected: 1},
		{input: "pwwkew", epected: 3},
		{input: "au", epected: 2},
	} {
		if test.epected != lengthOfLongestSubstring(test.input) {
			t.Errorf("expect %v but got %v at input %v", test.epected, lengthOfLongestSubstring(test.input), test.input)
		}
	}
}

// Runtime: 206 ms, faster than 23.15% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 6.8 MB, less than 22.59% of Go online submissions for Longest Substring Without Repeating Characters.
// Time: O(n^2)
// Space: O(n)

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) <= 1 {
// 		return len(s)
// 	}

// 	var longestSub string

// out:
// 	for i := 0; i < len(s); i++ {
// 		cache := string(s[i])
// 		for j := i + 1; j < len(s); j++ {
// 			if !strings.Contains(cache, string(s[j])) {
// 				cache += string(s[j])
// 				if len(longestSub) < len(cache) {
// 					longestSub = cache
// 				}
// 			} else {
// 				if len(longestSub) < len(cache) {
// 					longestSub = cache
// 				}
// 				continue out
// 			}
// 		}

// 	}

// 	return len(longestSub)
// }

// Runtime: 4 ms, faster than 91.69% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 3.1 MB, less than 58.46% of Go online submissions for Longest Substring Without Repeating Characters.

// Time O(n)
// Space O(n)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	cache, max, left := map[rune]int{}, 0, 0
	for right, char := range s {
		if _, ok := cache[char]; ok && cache[char] >= left {
			if right-left > max {
				max = right - left
			}
			left = cache[char] + 1
		}
		cache[char] = right
	}

	if len(s)-left > max {
		max = len(s) - left
	}

	return max
}
