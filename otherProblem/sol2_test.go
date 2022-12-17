package sol

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		message string
		K       int
		expect  string
	}{
		{message: "And now here is my secret", K: 15, expect: "And now ..."},
		{message: "There is an animal with four legs", K: 15, expect: "There is an ..."},
		{message: "super dog", K: 4, expect: "..."},
		{message: "", K: 4, expect: "..."},
		{message: "how are you", K: 20, expect: "how are you"},
	} {
		if Solution2(test.message, test.K) != test.expect {
			t.Errorf("expect [%v] but got [%v]", test.expect, Solution2(test.message, test.K))
		}
	}
}

func Solution2(message string, K int) string {
	// write your code in Go 1.13.8

	const (
		empty = "..."
	)
	leng := len(message)
	if leng <= K {
		return message
	}

	words := strings.Fields(message)

	if len(words) == 0 {
		return empty
	}

	res := words[0]
	if len(res) > K {
		return empty
	} else if len(res) == K {
		return words[0]
	}

	for i := 1; i < len(words); i++ {
		concate := fmt.Sprintf(res+" %v", words[i])
		if i < len(words)-2 {
			if len(concate) < K {
				concateWithEmpty := fmt.Sprintf(concate+" %v", empty)
				if len(concateWithEmpty) > K {
					return res + " " + empty
				} else {
					res = concate
				}
			} else {
				return res + " " + empty
			}
		} else {
			if len(concate) <= K {
				return concate
			} else {
				return res
			}
		}

	}

	return res
}
