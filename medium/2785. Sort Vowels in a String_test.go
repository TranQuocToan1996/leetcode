package medium

import (
	"strings"
	"testing"
)

// Time: O(n)
// Space: O(1) for vowelsCount only 10 vowels
// O(n) for the output
func sortVowels(s string) string {
	var (
		n           = len(s)
		sortedVowel = "AEIOUaeiou"
		res         strings.Builder
		vowelsCount = make(map[byte]int)
	)
	res.Grow(n)

	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'o' || c == 'u' || c == 'i' ||
			c == 'A' || c == 'E' || c == 'O' || c == 'U' || c == 'I'
	}

	for i := 0; i < n; i++ {
		if isVowel(s[i]) {
			vowelsCount[s[i]]++
		}
	}

	j := 0
	for i := 0; i < n; i++ {
		if !isVowel(s[i]) {
			res.WriteByte(s[i])
			continue
		}

		for vowelsCount[sortedVowel[j]] == 0 {
			j++
		}

		res.WriteByte(sortedVowel[j])
		vowelsCount[sortedVowel[j]]--
	}

	return res.String()
}

func TestSortVowels(t *testing.T) {
	for _, test := range []struct {
		input  string
		output string
	}{
		{input: "lEetcOde", output: "lEOtcede"},
		{input: "lYmpH", output: "lYmpH"},
	} {
		res := sortVowels(test.input)
		if test.output != res {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
