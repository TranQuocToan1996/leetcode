package easy

import "testing"

// No loop
// Time, space: O(n)
func destCity(paths [][]string) string {
	seen := map[string]string{}
	for _, path := range paths {
		from, to := path[0], path[1]
		seen[from] = to
	}

	for _, to := range seen {
		_, ok := seen[to]
		if !ok {
			return to
		}
	}

	return ""
}

func TestDestCity(t *testing.T) {
	for _, test := range []struct {
		paths  [][]string
		output string
	}{
		{paths: [][]string{
			{"London", "New York"},
			{"New York", "Lima"},
			{"Lima", "Sao Paulo"},
		}, output: "Sao Paulo"},
	} {
		res := destCity(test.paths)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
