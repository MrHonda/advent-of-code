package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	dir_up    = "^"
	dir_down  = "v"
	dir_left  = ">"
	dir_right = "<"
)

const block = "#"

var dirs = []string{dir_up, dir_down, dir_left, dir_right}

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(grid [][]string) {
	visited := []string{}
	maxI := len(grid) - 1
	maxJ := len(grid[0]) - 1
	i, j := getGridCurrentPos(grid)
	dir := grid[i][j]

loop:
	for {
		currentVisit := string(i) + "," + string(j)

		if !slices.Contains(visited, currentVisit) {
			visited = append(visited, currentVisit)
		}

		switch dir {
		case dir_up:
			newI := i - 1
			if newI < 0 {
				break loop
			}

			if grid[newI][j] == block {
				dir = dir_right
			} else {
				i = newI
			}
		case dir_down:
			newI := i + 1
			if newI > maxI {
				break loop
			}

			if grid[newI][j] == block {
				dir = dir_left
			} else {
				i = newI
			}
		case dir_left:
			newJ := j - 1
			if newJ < 0 {
				break loop
			}

			if grid[i][newJ] == block {
				dir = dir_up
			} else {
				j = newJ
			}
		case dir_right:
			newJ := j + 1
			if newJ > maxJ {
				break loop
			}

			if grid[i][newJ] == block {
				dir = dir_down
			} else {
				j = newJ
			}
		default:
			panic("unkown dir" + dir)
		}
	}

	fmt.Println("part1", len(visited))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][]string {
	lines := strings.Split(string(data), "\r\n")
	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func getGridCurrentPos(grid [][]string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if slices.Contains(dirs, grid[i][j]) {
				return i, j
			}

		}
	}

	panic("current position not found")
}
