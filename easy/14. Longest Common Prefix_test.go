package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	for _, test := range []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
	} {
		if longestCommonPrefix(test.strs) != test.expected {
			t.Error("fail at ", test.expected)
		}
	}
}

func longestCommonPrefix(strs []string) string {
	min := 201 // Test contrains is length 200
	for _, val := range strs {
		if len(val) < min {
			min = len(val)
		}
	}

loop:
	for i := min; i > 0; i-- {
		same := strs[0][:i]
		for j := 1; j < len(strs); j++ {
			sameInside := strs[j][:i]
			if same != sameInside {
				continue loop
			}
		}
		return same
	}
	return ""
}
