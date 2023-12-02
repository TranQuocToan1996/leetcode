package easy

import "testing"

// Time: O(n + k)
// Space: O(1)
func countCharacters(words []string, chars string) int {
	res := 0
	countChars := map[rune]int{}
	for _, char := range chars {
		countChars[char]++
	}

	var countWords map[rune]int
	for _, word := range words {
		countWords = map[rune]int{}
		for _, c := range word {
			countWords[c]++
		}
		match := true
		for c, count := range countWords {
			if countChars[c] < count {
				match = false
				break
			}
		}
		if match {
			res += len(word)
		}
	}

	return res
}

func TestCountCharacters(t *testing.T) {
	for _, test := range []struct {
		words  []string
		chars  string
		output int
	}{
		{words: []string{"cat", "bt", "hat", "tree"}, chars: "atach", output: 6},
		{words: []string{"hello", "world", "leetcode"}, chars: "welldonehoneyr", output: 10},
	} {
		res := countCharacters(test.words, test.chars)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
