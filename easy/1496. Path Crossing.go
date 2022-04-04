package main

import (
	"fmt"
)

func main() {
	// first checking
	paths := []string{
		"NESWW",
		"NES",
	}
	for _, path := range paths {
		fmt.Println(isPathCrossing(path))
	}
}

// https://leetcode.com/problems/path-crossing/
/*
time complexity: O(n) increses linear with the length of path
Space complexity: O(n) increses linear with the length of path
*/
func isPathCrossing(path string) bool {
	type coordinates struct {
		X int
		Y int
	}
	x := 0
	y := 0
	visited := make(map[coordinates]bool)
	// Set the first position with x = 0, y = 0
	visited[coordinates{x, y}] = true
	for _, direction := range path {
		// 'W' is a rune, "W" is a string
		switch direction {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		if visited[coordinates{x, y}] == true {
			return true
		}
		visited[coordinates{x, y}] = true
	}

	return false
}
