package main

//https://www.youtube.com/watch?v=layyhtQQuM0

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

// Queue represents a simple queue
type Queue struct {
	items []Position
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item Position) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() Position {
	if len(q.items) == 0 {
		panic("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(grid [][]int) {
	trailheads := findTrailheads(grid)
	result := 0

	for _, trailhead := range trailheads {
		result += score(grid, trailhead.row, trailhead.col)
	}

	fmt.Println("part1", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][]int {
	lines := strings.Split(string(data), "\r\n")
	grid := [][]int{}

	for _, line := range lines {
		vals := strings.Split(line, "")
		row := []int{}

		for _, val := range vals {
			intVal, err := strconv.Atoi(val)
			handleError(err)
			row = append(row, intVal)
		}

		grid = append(grid, row)
	}

	return grid
}

func findTrailheads(grid [][]int) []Position {
	trailheads := []Position{}

	for r, row := range grid {
		for c, col := range row {
			if col == 0 {
				trailheads = append(trailheads, Position{row: r, col: c})
			}
		}
	}

	return trailheads
}

func score(grid [][]int, r int, c int) int {
	rowsCount := len(grid)
	colsCount := len(grid[0])

	queue := Queue{}
	queue.Enqueue(Position{row: r, col: c})
	seen := []Position{{row: r, col: c}}
	summits := 0

	for len(queue.items) > 0 {
		currPos := queue.Dequeue()
		neighbors := []Position{
			{currPos.row - 1, currPos.col},
			{currPos.row, currPos.col + 1},
			{currPos.row + 1, currPos.col},
			{currPos.row, currPos.col - 1},
		}

		for _, neighbor := range neighbors {
			if neighbor.row < 0 || neighbor.col < 0 || neighbor.row >= rowsCount || neighbor.col >= colsCount {
				continue
			}
			if grid[neighbor.row][neighbor.col] != grid[currPos.row][currPos.col]+1 {
				continue
			}

			if slices.Contains(seen, neighbor) {
				continue
			}

			seen = append(seen, neighbor)

			if grid[neighbor.row][neighbor.col] == 9 {
				summits++
			} else {
				queue.Enqueue(neighbor)
			}
		}
	}

	return summits
}
