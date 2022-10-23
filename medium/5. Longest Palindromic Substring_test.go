package medium

import (
	"strings"
	"testing"
)

// https://leetcode.com/problems/longest-palindromic-substring/

func TestLongestPalindrome(t *testing.T) {
	for _, test := range []struct {
		input   string
		expect  string
		expect2 string
	}{
		{input: "babad", expect: "bab", expect2: "aba"},
		{input: "cbbd", expect: "bb", expect2: ""},
		{input: "bb", expect: "bb", expect2: ""},
		{input: "ac", expect: "a", expect2: "c"},
	} {
		if !strings.EqualFold(longestPalindrome(test.input), test.expect) {
			isErr := false
			if len(test.expect2) == 0 {
				isErr = true
			} else {
				if !strings.EqualFold(longestPalindrome(test.input), test.expect2) {
					isErr = true
				}
			}

			if isErr {
				t.Errorf("expect %v but got %v", test.expect, longestPalindrome(test.input))
			}
		}
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

// Runtime: 82 ms, faster than 35.23% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.5 MB, less than 66.57% of Go online submissions for Longest Palindromic Substring.

// Time: O(n^2)
// space: O(n)

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	max := len(s)
	res := ""
	for max > 1 {
		for i := 0; i+max-1 < len(s); i++ {
			sub := s[i : i+max]
			if isPalindrome(sub) {
				return sub
			}
		}
		max--
	}

	if max <= 1 {
		return s[0:1]
	}

	return res
}
