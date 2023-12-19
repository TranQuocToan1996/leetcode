package easy

import (
	"reflect"
	"testing"
)

const rangeRGB = 256

// Time: (n * m)
// Space: O(1)
func imageSmoother(img [][]int) [][]int {
	n, m := len(img), len(img[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			img[i][j] += encodeSum(img, i, j)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			img[i][j] /= rangeRGB
		}
	}

	return img
}

func encodeSum(img [][]int, i, j int) int {
	sum, count := 0, 0
	n, m := len(img), len(img[0])
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if 0 <= x && x < n && 0 <= y && y < m {
				sum += img[x][y] % rangeRGB
				count += 1
			}
		}
	}

	return (sum / count) * rangeRGB
}

func TestImageSmoother(t *testing.T) {
	for _, test := range []struct {
		img    [][]int
		output [][]int
	}{
		{img: [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, output: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}},
		{img: [][]int{
			{1, 1, 1},
			{1, 0, 1},
		}, output: [][]int{
			{0, 0, 0},
			{0, 0, 0},
		}},
		{img: [][]int{
			{100, 200, 100},
			{200, 50, 200},
			{100, 200, 100},
		}, output: [][]int{
			{137, 141, 137},
			{141, 138, 141},
			{137, 141, 137},
		}},
	} {
		res := imageSmoother(test.img)
		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
