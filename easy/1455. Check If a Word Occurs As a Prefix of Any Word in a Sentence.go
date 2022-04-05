package main

import (
	"fmt"
	"strings"
)

func main() {
	type word struct {
		sentence   string
		searchWord string
	}

	tests := []word{
		{"i love eating burger", "burg"},
		{"this problem is an easy problem", "pro"},
		{"i am tired", "you"},
	}
	for _, test := range tests {
		fmt.Println(isPrefixOfWord(test.sentence, test.searchWord))
	}
}

// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/

/*
time complexity: O(),  ms
Space complexity: O(),  1.9 MB
*/
func isPrefixOfWord(sentence string, searchWord string) int {
	arr := strings.Split(sentence, " ")
	for index, word := range arr {
		if strings.HasPrefix(word, searchWord) {
			return index + 1
		}
	}
	return -1
}
