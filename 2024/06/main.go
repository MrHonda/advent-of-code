package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
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
	part2(parseData(string(data)))
}

func part1(grid [][]string) {
	visited, _ := getVisited(grid)

	fmt.Println("part1", len(visited))
}

func part2(grid [][]string) {
	result := 0
	visited, _ := getVisited(grid)
	startI, startJ := getGridCurrentPos(grid)

	for i, visitVal := range visited {
		values := strings.Split(visitVal, ",")
		valI, err1 := strconv.Atoi(values[0])
		handleError(err1)

		valJ, err2 := strconv.Atoi(values[1])
		handleError(err2)

		if (startI == valI && startJ == valJ) || grid[valI][valJ] == block {
			continue
		}

		gridVal := grid[valI][valJ]
		grid[valI][valJ] = block
		_, loop := getVisited(grid)
		grid[valI][valJ] = gridVal

		if loop {
			result++
		}

		println(i, len(visited))

	}

	fmt.Println("part2", result)
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

func getVisited(grid [][]string) ([]string, bool) {
	visited := []string{}
	maxI := len(grid) - 1
	maxJ := len(grid[0]) - 1
	i, j := getGridCurrentPos(grid)
	dir := grid[i][j]
	tries := len(grid) * len(grid[0])
	try := 0
loop:
	for {
		if try >= tries {
			return nil, true
		}

		currentVisit := strconv.Itoa(i) + "," + strconv.Itoa(j)

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

		try++
		// println(try, tries)
	}

	return visited, false
}
