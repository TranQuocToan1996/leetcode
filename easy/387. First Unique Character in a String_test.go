package easy

import "testing"

func TestFirstUniQueChar(t *testing.T) {
	for _, test := range []struct {
		input  string
		output int
	}{
		{input: "leetcode", output: 0},
		{input: "loveleetcode", output: 2},
		{input: "aabb", output: -1},
	} {
		if test.output != firstUniqChar(test.input) {
			t.Error("fail")
		}
	}
}

type counts struct {
	count      int
	firstIndex int
}

func firstUniqChar(s string) int {

	var exist = map[byte]counts{}

	for i := 0; i < len(s); i++ {
		_, ok := exist[s[i]]
		if !ok {
			exist[s[i]] = counts{count: 1, firstIndex: i}
		} else {
			temp := exist[s[i]]
			temp.count++
			exist[s[i]] = temp
		}
	}

	for i := 0; i < len(s); i++ {
		if exist[s[i]].count == 1 {
			return exist[s[i]].firstIndex
		}
	}

	return -1
}
