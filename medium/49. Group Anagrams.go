package medium

import "sort"

func groupAnagrams(strs []string) [][]string {
	var (
		res  = [][]string{}
		seen = make(map[string][]string)
	)

	for _, str := range strs {
		sortString := sortString(str)
		seen[sortString] = append(seen[sortString], str)
	}

	for _, list := range seen {
		res = append(res, list)
	}

	return res
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
