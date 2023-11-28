package medium

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

type decodeStringData struct {
	num        int
	privString string
}

func decodeString(s string) string {
	var stack []decodeStringData
	curString := ""
	curNum := 0

	for _, c := range s {
		switch {
		case c == '[':
			stack = append(stack, decodeStringData{curNum, curString})
			curNum = 0
			curString = ""
		case c == ']':
			pop := stack[len(stack)-1]
			curString = pop.privString + multiplyString(curString, pop.num)
			stack = stack[:len(stack)-1]
		case unicode.IsDigit(c):
			curNum = curNum*10 + int(c-'0')
		default:
			curString += string(c)
		}
	}

	return curString
}

func multiplyString(s string, num int) string {
	builder := strings.Builder{}
	builder.Grow(len(s) * num)
	for i := 0; i < num; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}

func TestDecodeString(t *testing.T) {
	for _, test := range []struct {
		s      string
		output string
	}{
		{s: "3[a]2[bc]", output: "aaabcbc"},
		{s: "3[a2[c]]", output: "accaccacc"},
		{s: "2[abc]3[cd]ef", output: "abcabccdcdcdef"},
	} {
		res := decodeString(test.s)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
