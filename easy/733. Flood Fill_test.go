package easy

import (
	"reflect"
	"testing"
)

// Time O(m*n)
// Space O(m*n)
// func floodFill(image [][]int, sr int, sc int, color int) [][]int {
// 	type data struct {
// 		sr, sc int
// 	}
// 	startColor := image[sr][sc]
// 	m, n := len(image), len(image[0])
// 	queue := make([]data, 0, m*n)
// 	queue = append(queue, data{sr: sr, sc: sc})
// 	bfs := func(d data) {
// 		if d.sr >= m || d.sr < 0 || d.sc >= n || d.sc < 0 || startColor == color || image[d.sr][d.sc] != startColor {
// 			return
// 		}
// 		image[d.sr][d.sc] = color
// 		queue = append(queue, data{sr: d.sr + 1, sc: d.sc})
// 		queue = append(queue, data{sr: d.sr - 1, sc: d.sc})
// 		queue = append(queue, data{sr: d.sr, sc: d.sc + 1})
// 		queue = append(queue, data{sr: d.sr, sc: d.sc - 1})
// 	}
// 	for len(queue) > 0 {
// 		data := queue[0]
// 		queue = queue[1:]
// 		bfs(data)
// 	}
// 	return image
// }

// Time O(m*n)
// Space O(m*n)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startColor := image[sr][sc]
	m, n := len(image), len(image[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= m || x < 0 || y >= n || y < 0 || startColor == color || image[x][y] != startColor {
			return
		}
		image[x][y] = color
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	dfs(sr, sc)
	return image
}

func TestFloodFill(t *testing.T) {
	for _, test := range []struct {
		image, expect [][]int
		sr, sc        int
		color         int
	}{
		{
			sr:    1,
			sc:    1,
			color: 2,
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			}, expect: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			sr:    0,
			sc:    0,
			color: 0,
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			}, expect: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	} {
		res := floodFill(test.image, test.sr, test.sc, test.color)
		if !reflect.DeepEqual(res, test.expect) {
			t.Errorf("expect %v but got %v", test.expect, test.image)
		}
	}
}
