package easy

import (
	"testing"

	"leetcode/uti"
)

func minOperations(s string) int {
	n := len(s)
	start0 := 0
	for i := range s {
		if i%2 == 0 {
			if s[i] == '1' {
				start0++
			}
		} else {
			if s[i] == '0' {
				start0++
			}
		}
	}

	return uti.Min(start0, n-start0)
}

func TestMinOperations(t *testing.T) {
	for _, test := range []struct {
		s      string
		output int
	}{
		{s: "0100", output: 1},
		{s: "10", output: 0},
		{s: "1111", output: 2},
		{s: "110010", output: 2},
	} {
		res := minOperations(test.s)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
