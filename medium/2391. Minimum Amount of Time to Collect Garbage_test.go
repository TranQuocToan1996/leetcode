package medium

import (
	"strings"
	"testing"
)

func garbageCollection(garbage []string, travel []int) int {
	const (
		G = "G"
		M = "M"
		P = "P"
	)
	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i-1]
	}
	res := 0
	maxM, maxP, maxG := 0, 0, 0
	for i := len(garbage) - 1; i >= 0; i-- {
		res += len(garbage[i])
		if strings.Contains(garbage[i], G) && maxG < i {
			maxG = i
		}
		if strings.Contains(garbage[i], M) && maxM < i {
			maxM = i
		}
		if strings.Contains(garbage[i], P) && maxP < i {
			maxP = i
		}
	}
	res += travelTime(travel, maxM)
	res += travelTime(travel, maxP)
	res += travelTime(travel, maxG)
	return res
}

func travelTime(travel []int, maxIndex int) int {
	if maxIndex > 0 {
		return travel[maxIndex-1]
	}
	return 0
}

func TestGarbageCollection(t *testing.T) {
	for _, test := range []struct {
		garbage []string
		travel  []int
		output  int
	}{
		{garbage: []string{"G", "P", "GP", "GG"}, travel: []int{2, 4, 3}, output: 21},
		{garbage: []string{"MMM", "PGM", "GP"}, travel: []int{3, 10}, output: 37},
	} {
		res := garbageCollection(test.garbage, test.travel)
		if res != test.output {
			t.Errorf("expect %v but got %v", test.output, res)
		}
	}
}
