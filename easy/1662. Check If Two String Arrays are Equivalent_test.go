package easy

import (
	"testing"
)

// func arrayStringsAreEqual(word1 []string, word2 []string) bool {
// 	return strings.Join(word1, "") == strings.Join(word2, "")
// }

// Time: O(n * k)
// Space: O(1)
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var (
		n1, n2                         = len(word1), len(word2)
		wordPointer1, wordPointer2     = 0, 0
		stringPointer1, stringPointer2 = 0, 0
	)

	for wordPointer1 < n1 && wordPointer2 < n2 {
		if word1[wordPointer1][stringPointer1] !=
			word2[wordPointer2][stringPointer2] {
			return false
		}
		stringPointer1++
		stringPointer2++
		if stringPointer1 == len(word1[wordPointer1]) {
			wordPointer1++
			stringPointer1 = 0
		}
		if stringPointer2 == len(word2[wordPointer2]) {
			wordPointer2++
			stringPointer2 = 0
		}

	}

	return wordPointer1 == n1 && wordPointer2 == n2
}

func TestArrayStringsAreEqual(t *testing.T) {
	for _, test := range []struct {
		word1, word2 []string
		output       bool
	}{
		{word1: []string{"ab", "c"}, word2: []string{"a", "bc"}, output: true},
		{word1: []string{"a", "cb"}, word2: []string{"ab", "c"}, output: false},
		{word1: []string{"abc", "d", "defg"}, word2: []string{"abcddefg"}, output: true},
	} {
		res := arrayStringsAreEqual(test.word1, test.word2)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
