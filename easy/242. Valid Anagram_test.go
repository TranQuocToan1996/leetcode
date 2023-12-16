package easy

import "testing"

// Time: O(n)
// Space: O(1)
func isAnagram(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}
	anagramMap := [26]rune{}
	for _, c := range s {
		anagramMap[c-'a']++
	}
	for _, c := range p {
		if anagramMap[c-'a'] == 0 {
			return false
		}
		anagramMap[c-'a']--
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	for _, test := range []struct {
		s      string
		p      string
		output bool
	}{
		{s: "anagram", p: "nagaram", output: true},
		{s: "rat", p: "car", output: false},
	} {
		res := isAnagram(test.s, test.p)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
