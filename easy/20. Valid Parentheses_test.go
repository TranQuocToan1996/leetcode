package easy

import "testing"

// Time: O(n): loop through n char of string
// Space: O(n): memory for queue, maximum in case all chars are opening parentheses

// Runtime: 3 ms, faster than 36.08% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 43.16% of Go online submissions for Valid Parentheses.
func TestIsValid(t *testing.T) {
	for _, test := range []struct {
		input    string
		expected bool
	}{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
	} {
		if isValid(test.input) != test.expected {
			t.Errorf("%v: Expect %v but actually got %v", test.input, test.expected, isValid(test.input))
		}
	}

}

func matchOpenParentheses(char rune) bool {
	list := `([{`
	for _, parenthes := range list {
		if char == parenthes {
			return true
		}
	}

	return false
}

func matchCloseParentheses(char rune) bool {
	list := `)]}`
	for _, parenthes := range list {
		if char == parenthes {
			return true
		}
	}

	return false
}

func parse(open rune) (close rune) {
	switch open {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	}

	return 0
}

func isValid(s string) bool {

	queue := []rune{}

	for _, char := range s {
		if matchOpenParentheses(char) {
			queue = append(queue, char)
			continue
		}

		if matchCloseParentheses(char) {
			if len(queue) == 0 {
				return false
			}

			parse := parse(queue[len(queue)-1])
			if parse == 0 {
				return false
			}

			if char != parse {
				return false
			} else {
				queue = queue[:len(queue)-1]
			}
		}
	}

	return len(queue) == 0
}
