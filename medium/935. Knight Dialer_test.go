package medium

import "testing"

// Time: out of time (3^n)
// func knightDialer(n int) int {
// 	if n == 1 {
// 		return 10
// 	}
// 	const mod = int(1e9 + 7)
// 	// first move
// 	var (
// 		possibleMoves = map[int][]int{
// 			0: {4, 6},
// 			1: {6, 8},
// 			2: {7, 9},
// 			3: {4, 8},
// 			4: {0, 3, 9},
// 			5: nil,
// 			6: {0, 1, 7},
// 			7: {2, 6},
// 			8: {1, 3},
// 			9: {2, 4},
// 		}
// 		curPos = map[int]int{
// 			0: 1,
// 			1: 1,
// 			2: 1,
// 			3: 1,
// 			4: 1,
// 			5: 1,
// 			6: 1,
// 			7: 1,
// 			8: 1,
// 			9: 1,
// 		}
// 		res = 0
// 	)
// 	n -= 2
// 	for n > 0 {
// 		next := map[int]int{}
// 		for pos, count := range curPos {
// 			possibleMove := possibleMoves[pos]
// 			for k := range possibleMove {
// 				next[possibleMove[k]] += count
// 			}
// 		}

// 		curPos = next
// 		n--
// 	}
// 	for nextPos, count := range curPos {
// 		possibleMove := possibleMoves[nextPos]
// 		res = (res + count*len(possibleMove)) % mod
// 	}
// 	return res % mod
// }

// Time: O(n)
// Space: O(n)
// func knightDialer(n int) int {
// 	const mod = int(1e9 + 7)
// 	memo := make([][10]int, n)

// 	var topDownDP func(remain, curPos int) int
// 	topDownDP = func(remain, curPos int) int {
// 		if remain == 0 {
// 			return 1
// 		}
// 		if curPos == 5 {
// 			return 0
// 		}
// 		val := memo[remain][curPos]
// 		if val != 0 {
// 			return val
// 		}
// 		res := 0
// 		for _, nextPos := range moves[curPos] {
// 			res = (res + topDownDP(remain-1, nextPos)) % mod
// 		}
// 		memo[remain][curPos] = res
// 		return res
// 	}
// 	res := 0
// 	for i := 0; i < 10; i++ {
// 		res = (res + topDownDP(n-1, i)) % mod
// 	}
// 	return res
// }

var moves = [][]int{
	{4, 6},
	{6, 8},
	{7, 9},
	{4, 8},
	{3, 9, 0},
	{}, // pos 5 can not move
	{1, 7, 0},
	{2, 6},
	{1, 3},
	{2, 4},
}

// Time: O(n)
// Space: O(n)
// func knightDialer(n int) int {
// 	const mod = int(1e9 + 7)
// 	memo := make([][10]int, n)
// 	for k := range memo[0] {
// 		memo[0][k] = 1
// 	}
// 	m := len(memo[0])
// 	for remain := 1; remain < n; remain++ {
// 		for pos := 0; pos < m; pos++ {
// 			res := 0
// 			for _, nextPos := range moves[pos] {
// 				res = (res + memo[remain-1][nextPos]) % mod
// 			}
// 			memo[remain][pos] = res
// 		}
// 	}
// 	res := 0
// 	for pos := 0; pos < m; pos++ {
// 		res = (res + memo[n-1][pos]) % mod
// 	}

// 	return res
// }

// Time: O(n)
// Space: O(1)
func knightDialer(n int) int {
	const mod = int(1e9 + 7)
	var priv, cur [10]int
	for k := range priv {
		priv[k] = 1
	}
	for remain := 1; remain < n; remain++ {
		cur = [10]int{}
		for pos := 0; pos < len(priv); pos++ {
			for _, nextPos := range moves[pos] {
				cur[pos] = (cur[pos] + priv[nextPos]) % mod
			}
		}

		priv = cur
	}

	res := 0
	for pos := 0; pos < len(priv); pos++ {
		res = (res + priv[pos]) % mod
	}

	return res
}

func TestKnightDialer(t *testing.T) {
	for _, test := range []struct {
		n      int
		output int
	}{
		{n: 3131, output: 136006598}, // Mod test
		{n: 1, output: 10},
		{n: 2, output: 20},
		{n: 3, output: 46},
		{n: 4, output: 104},
	} {
		res := knightDialer(test.n)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
