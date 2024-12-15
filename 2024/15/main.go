package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	dir_up    = "^"
	dir_down  = "v"
	dir_left  = "<"
	dir_right = ">"
)

const (
	char_wall  = "#"
	char_box   = "O"
	char_free  = "."
	char_robot = "@"
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(grid [][]string, moves []string) {
	robotPos := findRobotPos(grid)

	// printGrid(grid)
	// fmt.Println("=================")

	for _, move := range moves {
		robotPos = moveRobot(grid, robotPos, getMoveDir(move))
		// printGrid(grid)
		// fmt.Println("=================")
	}

	result := 0

	for r, row := range grid {
		for c, char := range row {
			if char == char_box {
				result += 100*r + c
			}
		}
	}

	fmt.Println("part1", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) ([][]string, []string) {
	parts := strings.Split(data, "\r\n\r\n")
	grid := [][]string{}

	lines := strings.Split(parts[0], "\r\n")

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	moveLines := strings.Split(parts[1], "\r\n")
	moves := []string{}

	for _, line := range moveLines {
		moves = append(moves, strings.Split(line, "")...)
	}

	return grid, moves
}

func findRobotPos(grid [][]string) [2]int {
	for r, row := range grid {
		for c, char := range row {
			if char == char_robot {
				return [2]int{r, c}
			}
		}
	}

	panic("robot pos not found")
}

func getMoveDir(move string) [2]int {
	if move == dir_left {
		return [2]int{0, -1}
	}
	if move == dir_right {
		return [2]int{0, 1}
	}
	if move == dir_up {
		return [2]int{-1, 0}
	}
	if move == dir_down {
		return [2]int{1, 0}
	}

	panic("invalid move" + move)
}

func moveRobot(grid [][]string, robotPos [2]int, moveDir [2]int) [2]int {
	newRobotPos := [2]int{robotPos[0] + moveDir[0], robotPos[1] + moveDir[1]}
	char := grid[newRobotPos[0]][newRobotPos[1]]

	if char == char_wall {
		return robotPos
	}

	if char == char_free {
		grid[robotPos[0]][robotPos[1]] = char_free
		grid[newRobotPos[0]][newRobotPos[1]] = char_robot
		return newRobotPos
	}

	boxes := [][2]int{}
	pos := [2]int{robotPos[0], robotPos[1]}

	for {
		pos[0] += moveDir[0]
		pos[1] += moveDir[1]

		char := grid[pos[0]][pos[1]]

		if char == char_box {
			boxes = append(boxes, pos)
		} else if char == char_free {
			break
		} else if char == char_wall {
			boxes = nil
			break
		} else {
			panic("invalid char" + char)
		}
	}

	if len(boxes) == 0 {
		return robotPos
	}

	grid[robotPos[0]][robotPos[1]] = char_free

	for i, box := range boxes {
		if i == 0 {
			grid[box[0]][box[1]] = char_robot
		}

		grid[box[0]+moveDir[0]][box[1]+moveDir[1]] = char_box
	}

	return newRobotPos
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
