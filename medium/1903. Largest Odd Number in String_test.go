package medium

import "testing"

func largestOddNumber(num string) string {
	isOdd := func(num string) bool {
		return num == "1" || num == "3" || num == "5" || num == "7" || num == "9"
	}
	for i := len(num) - 1; i >= 0; i-- {
		if isOdd(string(num[i])) {
			return num[:i+1]
		}
	}
	return ""
}

func TestLargestOddNumber(t *testing.T) {
	for _, test := range []struct {
		num    string
		output string
	}{
		{num: "52", output: "5"},
		{num: "4206", output: ""},
		{num: "35427", output: "35427"},
	} {
		res := largestOddNumber(test.num)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
