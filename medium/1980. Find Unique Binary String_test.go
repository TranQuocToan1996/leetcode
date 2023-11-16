package medium

import (
	"testing"
)

// Time: O(n^2)
// Space: O(n)
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	match := map[string]bool{}
	for _, num := range nums {
		match[num] = true
	}

	var generate func(s string) string

	generate = func(s string) string {
		if len(s) == n {
			if !match[s] {
				return s
			}
			return ""
		}

		generateZero := generate(s + "0")
		if len(generateZero) > 0 {
			return generateZero
		}

		return generate(s + "1")
	}

	return generate("")
}

// func findDifferentBinaryString(nums []string) string {
// 	n := len(nums[0])
// 	match := map[int64]bool{}
// 	for _, num := range nums {
// 		match[binaryStringsToInt(num)] = true
// 	}

// 	begin := binaryStringsToInt(getBeginBinary(n))
// 	end := binaryStringsToInt(getEndBinary(n))

// 	for i := begin; i <= end; i++ {
// 		if !match[i] {
// 			return intToBinaryString(i, n)
// 		}
// 	}

// 	return ""
// }

// func getBeginBinary(n int) string {
// 	var builder strings.Builder
// 	for n > 0 {
// 		builder.WriteString("0")
// 		n--
// 	}
// 	return builder.String()
// }

// func getEndBinary(n int) string {
// 	var builder strings.Builder
// 	for n > 0 {
// 		builder.WriteString("1")
// 		n--
// 	}
// 	return builder.String()
// }

// func binaryStringsToInt(binary string) int64 {
// 	i, _ := strconv.ParseInt(binary, 2, 64)
// 	return i
// }

// func intToBinaryString(i int64, length int) string {
// 	binary := strconv.FormatInt(i, 2)
// 	for len(binary) < length {
// 		binary = "0" + binary
// 	}
// 	return binary
// }

func TestFindDifferentBinaryString(t *testing.T) {
	for _, test := range []struct {
		input  []string
		expect map[string]bool
	}{
		{input: []string{"0"}, expect: map[string]bool{
			"1": true,
		}},
		{input: []string{"01", "10"}, expect: map[string]bool{
			"11": true,
			"00": true,
		}},
		{input: []string{"00", "01"}, expect: map[string]bool{
			"11": true,
			"10": true,
		}},
		{input: []string{"111", "011", "001"}, expect: map[string]bool{
			"000": true,
			"010": true,
			"100": true,
			"110": true,
		}},
	} {
		res := findDifferentBinaryString(test.input)
		if !test.expect[res] {
			t.Errorf("expect any of key in map %v but got %v, input %v", test.expect, res, test.input)
		}
	}
}
