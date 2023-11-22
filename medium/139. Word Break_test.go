package medium

import (
	"testing"
	"time"

	"leetcode/uti"
)

// Time: O(n^3) n is len s, and m is len wordDict
// Space: O(n)
// DFS
// func wordBreak(s string, wordDict []string) bool {
// 	sort.Slice(wordDict, func(i, j int) bool {
// 		return len(wordDict[i]) > len(wordDict[j])
// 	})
// 	seen := make(map[string]bool)
// 	var dfs func(s string) bool
// 	dfs = func(s string) bool {
// 		if len(s) == 0 {
// 			return true
// 		}
// 		if seen[s] {
// 			return false
// 		}
// 		seen[s] = true
// 		for _, word := range wordDict {
// 			if strings.HasPrefix(s, word) {
// 				ok := dfs(s[len(word):])
// 				if ok {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(s)
// }

// Time: O(n*m) n is len s, and m is len wordDict
// Space: O(n)
// Approach 2: Top-Down Dynamic Programming
// func wordBreak(s string, wordDict []string) bool {
// 	memo := make([]int, len(s))
// 	for k := range memo {
// 		memo[k] = -1
// 	}

// 	var dp func(int) bool
// 	dp = func(i int) bool {
// 		if i < 0 {
// 			return true
// 		}

// 		if memo[i] != -1 {
// 			return memo[i] == 1
// 		}

// 		for _, word := range wordDict {
// 			if s[i-len(word)+1:i+1] == word && dp(i-len(word)) {
// 				memo[i] = 1
// 				return true
// 			}
// 		}

// 		memo[i] = 0
// 		return false
// 	}
// 	return dp(len(s) - 1)
// }

// Time: O(n*m) n is len s, and m is len wordDict
// Space: O(n)
// Approach 3: Bottom-Up Dynamic Programming
func wordBreak(s string, wordDict []string) bool {
	memo := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if i-len(word)+1 < 0 {
				continue
			}
			if i-len(word)+1 == 0 || memo[i-len(word)] {
				if s[i-len(word)+1:i+1] == word {
					memo[i] = true
					break
				}
			}
		}
	}

	return memo[len(s)-1]
}

func TestWordBreak(t *testing.T) {
	uti.TimeOut(time.Second * 5)
	for _, test := range []struct {
		s        string
		wordDict []string
		output   bool
	}{
		{s: "leetcode", wordDict: []string{"leet", "code"}, output: true},
		{s: "applepenapple", wordDict: []string{"apple", "pen"}, output: true},
		{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, output: false},
		{
			// Time complexity test
			s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, output: false,
		},
	} {
		res := wordBreak(test.s, test.wordDict)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
