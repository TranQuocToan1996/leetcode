package easy

func TestPowerOf4(t *testing.T) {
	isPowerOfFour(4)
}

func isPowerOfFour(n int) bool {

	for n/4 >= 4 {
		n = n / 4
	}

	if n <= 1 {
		return true
	} else {
		return false
	}

}
