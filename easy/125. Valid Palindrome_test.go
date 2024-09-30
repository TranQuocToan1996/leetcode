package easy

import (
	"strings"
	"testing"
	"unicode"
)

func TestIsPalindromeString(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{input: "A man, a plan, a canal: Panama", expected: true},
		{input: "race a car", expected: false},
		{input: " ", expected: true},
		{input: ".,", expected: true},
	} {
		if isPalindromeString(test.input) != test.expected {
			t.Errorf("%v: Expect %v but actually got %v", test.input, test.expected, isPalindromeString(test.input))
		}
	}
}

func isPalindromeString(s string) bool {
	toLowerAndAlphanumeric := func(s string) string {
		var builder strings.Builder
		for _, char := range strings.ToLower(s) {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				builder.WriteRune(char)
			}
		}
		return builder.String()
	}
	s = toLowerAndAlphanumeric(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
