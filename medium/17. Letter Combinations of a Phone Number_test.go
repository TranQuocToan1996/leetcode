package medium

import (
	"leetcode/uti"
	"testing"
	"time"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func TestLetterCombinations(t *testing.T) {
	defer uti.LogRuntime(time.Now())
	for _, test := range []struct {
		input  string
		output []string
	}{
		{input: "23", output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{input: "", output: []string{}},
		{input: "2", output: []string{"a", "b", "c"}},
	} {
		result := letterCombinations(test.input)
		if !unorderMatch(result, test.output) {
			t.Errorf("Actually %v Expect %v", result, test.output)
		}
	}
}

func unorderMatch(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	m1, m2 := map[string]bool{}, map[string]bool{}
	for _, c := range arr1 {
		m1[c] = true
	}
	for _, c := range arr2 {
		m2[c] = true
	}
	for _, c := range arr1 {
		_, ok := m2[c]
		if !ok {
			return false
		}
	}
	for _, c := range arr2 {
		_, ok := m1[c]
		if !ok {
			return false
		}
	}
	return true
}

// The n is small so both cases doesn't product big difference
func letterCombinations(digits string) []string {

	// uti.LogRuntime ~ 7000 nsec, the iterator faster a bit in the first digit because we assign an array directly
	// Leetcode runtime 1 ms (61.52%)
	// Memery: 1.9 MB (88.94%)
	// Time complexity: for the result is (4^n*n) where 4 is case 7 or 9 that repeats 4 times
	// Space complexity: for the result is (4^n) where 4 is case 7 or 9
	return letterCombinationIteration(digits)

	// uti.LogRuntime ~ 8000 nsec
	// Leetcode runtime 1 ms (61.52%)
	// Memery: 1.9 MB (88.94%)
	// Time complexity: for the result is (4^n*n) where 4 is case 7 or 9 that repeats 4 times
	// Space complexity: Worst for callstack is O(n), for the result is (4^n) where 4 is case 7 or 9
	// return letterCombinationRecusive(digits)
}

var (
	dictionary = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
)

func letterCombinationIteration(digits string) []string {
	result := []string{}
	for i := range digits {
		chars := dictionary[digits[i]]
		if i == 0 {
			result = chars
			continue
		}
		newCombine := []string{}
		for _, char := range chars {
			for _, oldCombine := range result {
				newCombine = append(newCombine, oldCombine+char)
			}
		}
		result = newCombine
	}

	return result
}

func letterCombinationRecusive(digits string) []string {

	result := []string{}
	generateCombination("", digits, &result)

	return result
}

func generateCombination(current string, digits string, result *[]string) {
	if len(digits) == 0 {
		if current != "" {
			*result = append(*result, current)
		}
		return
	}

	currentDigit := digits[0]
	for _, char := range dictionary[currentDigit] {
		generateCombination(current+char, digits[1:], result)
	}
}