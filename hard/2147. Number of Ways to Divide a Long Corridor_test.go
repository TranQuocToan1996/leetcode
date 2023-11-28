package hard

import (
	"testing"
)

func numberOfWays(corridor string) int {
	const (
		mod = 1e9 + 7
		P   = 'P'
		S   = 'S'
	)
	seats, plants, res := 0, 0, 1
	for _, place := range corridor {
		if place == S {
			if seats == 2 {
				res = (res * (plants + 1)) % mod
				seats, plants = 0, 0
			}
			seats++
		} else if seats == 2 {
			plants++
		}
	}
	if seats == 2 {
		return res
	}
	return 0
}

func TestMaxProduct(t *testing.T) {
	for _, test := range []struct {
		corridor string
		output   int
	}{
		{corridor: "SSPPSPS", output: 3},
		{corridor: "PPSPSP", output: 1},
		{corridor: "S", output: 0},
		{corridor: "P", output: 0},
		{corridor: "PPPPPSPPSPPSPPPSPPPPSPPPPSPPPPSPPSPPPSPSPPPSPSPPPSPSPPPSPSPPPPSPPPPSPPPSPPSPPPPSPSPPPPSPSPPPPSPSPPPSPPSPPPPSPSPSS", output: 919999993},
	} {
		res := numberOfWays(test.corridor)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
