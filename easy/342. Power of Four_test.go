package easy

import "testing"

func TestPowerOf4(t *testing.T) {
	if !isPowerOfFour(4) || !isPowerOfFour(1) ||
		isPowerOfFour(-4) || isPowerOfFour(0) {
		t.Fatal("Fail")
	}
}

// Space O1
// Total Time log2(n)

func isPowerOfFour(n int) bool {

	// Time log2(n)
	for n >= 4 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}

	// Time O1
	if n == 1 {
		return true
	} else {
		return false
	}

}
