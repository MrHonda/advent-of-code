package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	r int //row
	c int //col
}

// Queue represents a simple queue
type Queue struct {
	items []Coord
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item Coord) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() Coord {
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
	part2(parseData(string(data)))
}

func part1(grid [][]string) {
	visited := map[Coord]bool{}
	result := 0

	for r := range grid {
		for c := range grid[r] {
			coord := Coord{r: r, c: c}

			if visited[coord] {
				continue
			}

			result += calculatePrice(findRegion(coord, visited, grid))
		}
	}

	fmt.Println("part1", result)
}

func part2(grid [][]string) {
	visited := map[Coord]bool{}
	result := 0

	for r := range grid {
		for c := range grid[r] {
			coord := Coord{r: r, c: c}

			if visited[coord] {
				continue
			}

			result += calculatePrice2(findRegion(coord, visited, grid))
		}
	}

	fmt.Println("part2", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][]string {
	grid := [][]string{}
	lines := strings.Split(data, "\r\n")

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func findRegion(startCoord Coord, visited map[Coord]bool, grid [][]string) (string, map[Coord]bool) {
	queue := Queue{items: []Coord{startCoord}}
	region := map[Coord]bool{}
	name := grid[startCoord.r][startCoord.c]

	for len(queue.items) > 0 {
		curCoord := queue.Dequeue()

		if visited[curCoord] {
			continue
		}

		currName := grid[curCoord.r][curCoord.c]

		if currName != name {
			continue
		}

		visited[curCoord] = true
		region[curCoord] = true

		//left
		if curCoord.c-1 >= 0 {
			queue.Enqueue(Coord{r: curCoord.r, c: curCoord.c - 1})
		}

		//top
		if curCoord.r-1 >= 0 {
			queue.Enqueue(Coord{r: curCoord.r - 1, c: curCoord.c})
		}

		//right
		if curCoord.c+1 < len(grid[curCoord.r]) {
			queue.Enqueue(Coord{r: curCoord.r, c: curCoord.c + 1})
		}

		//bottom
		if curCoord.r+1 < len(grid) {
			queue.Enqueue(Coord{r: curCoord.r + 1, c: curCoord.c})
		}
	}

	return name, region
}

func calculatePrice(name string, region map[Coord]bool) int {
	sides := 0

	for coord := range region {
		//left
		if !region[Coord{r: coord.r, c: coord.c - 1}] {
			sides++
		}

		//top
		if !region[Coord{r: coord.r - 1, c: coord.c}] {
			sides++
		}

		//right
		if !region[Coord{r: coord.r, c: coord.c + 1}] {
			sides++
		}

		//bottom
		if !region[Coord{r: coord.r + 1, c: coord.c}] {
			sides++
		}
	}

	price := len(region) * sides

	// fmt.Println("region:", name, "price:", price, region)

	return price
}

func findEdges(region map[Coord]bool) map[string][]Coord {
	edges := map[string][]Coord{}

	for coord := range region {
		//left
		if !region[Coord{r: coord.r, c: coord.c - 1}] {
			edges["L"] = append(edges["L"], coord)
		}

		//top
		if !region[Coord{r: coord.r - 1, c: coord.c}] {
			edges["T"] = append(edges["T"], coord)
		}

		//right
		if !region[Coord{r: coord.r, c: coord.c + 1}] {
			edges["R"] = append(edges["R"], coord)
		}

		//bottom
		if !region[Coord{r: coord.r + 1, c: coord.c}] {
			edges["B"] = append(edges["B"], coord)
		}
	}

	return edges
}

func calculatePrice2(name string, region map[Coord]bool) int {
	sides := 0
	edges := findEdges(region)

	sides += countVerticalSides(edges, "L")
	sides += countVerticalSides(edges, "R")
	sides += countHorizontalSides(edges, "T")
	sides += countHorizontalSides(edges, "B")

	price := len(region) * sides

	// fmt.Println("region:", name, "price:", price, edges)

	return price
}

func insertToSorted[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func countVerticalSides(edges map[string][]Coord, dir string) int {
	sides := 0
	colsMap := map[int][]int{}
	for _, coords := range edges[dir] {
		colsMap[coords.c] = insertToSorted(colsMap[coords.c], coords.r)
	}

	for _, rows := range colsMap {
		prevRow := rows[0]

		for i, row := range rows {
			if i == 0 {
				sides++
				continue
			}

			if row != prevRow+1 {
				sides++
			}

			prevRow = row
		}

	}

	return sides
}

func countHorizontalSides(edges map[string][]Coord, dir string) int {
	sides := 0
	rowsMap := map[int][]int{}
	for _, coords := range edges[dir] {
		rowsMap[coords.r] = insertToSorted(rowsMap[coords.r], coords.c)
	}

	for _, cols := range rowsMap {
		prevCol := cols[0]

		for i, col := range cols {
			if i == 0 {
				sides++
				continue
			}

			if col != prevCol+1 {
				sides++
			}

			prevCol = col
		}

	}

	return sides
}
