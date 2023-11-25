package medium

import "testing"

// Topological Sort | Kahn's Algorithm
// Time: O(numCourses + len(prerequisites))
// Space: O(n)
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		done     = 0
		graph    = make(map[int][]int, numCourses)
		inDegree = make([]int, numCourses)
		queue    []int
	)
	for _, prerequisite := range prerequisites {
		after := prerequisite[0]
		before := prerequisite[1]
		graph[before] = append(graph[before], after)
		inDegree[after] += 1
	}
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}
	for len(queue) > 0 {
		before := queue[0]
		queue = queue[1:]
		// Get list course learn after
		afterCourses := graph[before]

		// decrease inDegree of vertex point to
		for _, after := range afterCourses {
			inDegree[after]--
			if inDegree[after] == 0 {
				queue = append(queue, after)
			}
		}

		done++
	}
	return done == numCourses
}

func TestCanFinish(t *testing.T) {
	for _, test := range []struct {
		numCourses    int
		prerequisites [][]int
		output        bool
	}{
		{
			numCourses: 3,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
			},
			output: true,
		},
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0}, // To take course 1 you should have finished course 0.
			},
			output: true,
		},
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			output: false,
		},
	} {
		res := canFinish(test.numCourses, test.prerequisites)
		if res != test.output {
			t.Errorf("expect %v but got %v, test %v", test.output, res, test)
		}
	}
}
