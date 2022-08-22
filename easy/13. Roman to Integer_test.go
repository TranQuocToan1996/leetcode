package easy

import (
	"testing"
)

// https://leetcode.com/problems/roman-to-integer/

func TestRoman(t *testing.T) {
	if romanToInt("III") != 3 || romanToInt("LVIII") != 58 ||
		romanToInt("MCMXCIV") != 1994 {
		t.Fatal("fail")
	}
}

// // Space: Om
// // Time: On
func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && dict[s[i]] < dict[s[i+1]] {
			total -= dict[s[i]]
		} else {
			total += dict[s[i]]
		}
	}

	return total
}

// var (
// 	individualRoman = map[string]int{
// 		"I": 1,
// 		"V": 5,
// 		"X": 10,
// 		"L": 50,
// 		"C": 100,
// 		"D": 500,
// 		"M": 1000,
// 	}

// 	placeBefore = map[string]int{
// 		"IV": 4,
// 		"IX": 9,
// 		"XL": 40,
// 		"XC": 90,
// 		"CD": 400,
// 		"CM": 900,
// 	}
// )

// // Space: Om
// // Time: On x Om
// func romanToInt(s string) int {
// 	total := 0
// forLoop:
// 	for len(s) > 0 {

// 		for k, v := range placeBefore {
// 			if strings.HasPrefix(s, k) {
// 				total += v
// 				if len(s) > 2 {
// 					s = s[2:]
// 				} else {
// 					s = ""
// 				}
// 				continue forLoop
// 			}
// 		}

// 		for k, v := range individualRoman {
// 			if strings.HasPrefix(s, k) {
// 				total += v
// 				if len(s) > 1 {
// 					s = s[1:]
// 				} else {
// 					s = ""
// 				}
// 				continue forLoop
// 			}
// 		}

// 		// if no match return 0, maybe it is not a roman number
// 		return 0
// 	}

// 	return total
// }
