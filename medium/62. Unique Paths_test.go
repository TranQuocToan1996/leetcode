package medium

import (
	"reflect"
	"testing"
)

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 0
	}

	//

	return 0
}

func TestUniquePaths(t *testing.T) {
	for _, test := range []struct {
		m      int
		n      int
		output int
	}{
		{
			m:      3,
			n:      7,
			output: 28,
		},
		{
			m:      3,
			n:      2,
			output: 3,
		},
	} {
		res := uniquePaths(test.m, test.n)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %+v", test.output, res, test)
		}
	}
}
