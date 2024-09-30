package easy

import "testing"

func TestIsValid(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "[", expected: false},
		{input: "]", expected: false},
	} {
		if isValid(test.input) != test.expected {
			t.Errorf("%v: Expect %v but actually got %v", test.input, test.expected, isValid(test.input))
		}
	}
}

func isValid(s string) bool {
	getCloseParentheses := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	isOpen := map[rune]bool{
		'(': true,
		'{': true,
		'[': true,
	}
	var stack []rune
	for _, parenthese := range s {
		if isOpen[parenthese] {
			stack = append(stack, parenthese)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		// Pop recent parenthese from stack
		lastOpen := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if getCloseParentheses[lastOpen] != parenthese {
			return false
		}
	}
	return len(stack) == 0
}
