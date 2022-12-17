package sol

func Solution1(S string) int {
	leng := len(S)
	if leng <= 1 {
		return 0
	}

	countChars := map[rune]int{}
	for _, char := range S {
		countChars[char]++
	}

	odd := 0
	for _, count := range countChars {
		if count%2 != 0 {
			odd++
		}
	}

	if leng%2 == 0 {
		return odd
	} else {
		return odd - 1
	}

}
