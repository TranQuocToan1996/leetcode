package medium

import (
	"reflect"
	"testing"
)

// Out of time
// func findAnagrams(s string, p string) []int {
// 	var res []int
// 	countP := map[rune]int{}
// 	for _, c := range p {
// 		countP[c]++
// 	}

// 	isAnagram := func(s string) bool {
// 		lenS := len(s)
// 		lenP := len(p)

// 		if lenS != lenP {
// 			return false
// 		}

// 		anagramMap := make(map[string]int)

// 		for i := 0; i < lenS; i++ {
// 			anagramMap[string(s[i])]++
// 		}

// 		for i := 0; i < lenP; i++ {
// 			anagramMap[string(p[i])]--
// 		}

// 		for i := 0; i < lenS; i++ {
// 			if anagramMap[string(s[i])] != 0 {
// 				return false
// 			}
// 		}

// 		return true
// 	}

// 	for i := 0; i < len(s)-len(p)+1; i++ {
// 		if isAnagram(s[i : i+len(p)]) {
// 			res = append(res, i)
// 		}
// 	}

// 	return res
// }

// Time: O(n + k)
// Space: O(1) skip the space result
func findAnagrams(s string, p string) []int {
	var res []int
	countP := map[byte]int{}
	for i := 0; i < len(p); i++ {
		countP[p[i]]++
	}

	left, right := 0, 0
	for right < len(s) {
		count, ok := countP[s[right]]
		if !ok {
			for left < right {
				countP[s[left]]++
				left++
			}
			left++
			right++
		} else if count == 0 {
			countP[s[left]]++
			left++
		} else {
			countP[s[right]]--
			if (right-left+1 == len(p)) && countP[s[right]] == 0 {
				res = append(res, left)
			}
			right++
		}
	}

	return res
}

func TestFindAnagrams(t *testing.T) {
	for _, test := range []struct {
		s      string
		p      string
		output []int
	}{
		{s: "cbaebabacd", p: "abc", output: []int{0, 6}},
		{s: "abab", p: "ab", output: []int{0, 1, 2}},
		{s: "baa", p: "aa", output: []int{1}},
	} {
		res := findAnagrams(test.s, test.p)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}

func isAnagram(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)

	if lenS != lenP {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < lenP; i++ {
		anagramMap[string(p[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}
