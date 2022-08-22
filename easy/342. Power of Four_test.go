package easy

import "testing"

func TestPowerOf4(t *testing.T) {
	if !isPowerOfFour(4) || !isPowerOfFour(1) ||
		isPowerOfFour(-4) || isPowerOfFour(0) {
		t.Fatal("Fail")
	}
}

func isPowerOfFour(n int) bool {

	for n >= 4 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}

	if n == 1 {
		return true
	} else {
		return false
	}

}
