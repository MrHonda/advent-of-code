package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(data))
	part2(parseData(data))
}

func part1(data []string) {
	horizontalCount := countAll(data)
	verticalCount := countAll(getVertical(data))
	diagonalCount := countAll(getDiagonal(data))
	result := horizontalCount + verticalCount + diagonalCount

	fmt.Println("part1:", result, "| horizontal:", horizontalCount, "| vertical:", verticalCount, "| diagonal:", diagonalCount)
}

func part2(data []string) {
	result := 0
	grid := make([][]string, len(data))

	for i, line := range data {
		grid[i] = strings.Split(line, "")
	}

	for i, row := range grid {
		for j, col := range row {
			if col != "A" {
				continue
			}

			up := i - 1
			down := i + 1
			left := j - 1
			right := j + 1

			if up < 0 || left < 0 || down >= len(grid) || right >= len(row) {
				continue
			}

			if ((grid[up][left] == "M" && grid[down][right] == "S") || (grid[up][left] == "S" && grid[down][right] == "M")) &&
				((grid[up][right] == "M" && grid[down][left] == "S") || (grid[up][right] == "S" && grid[down][left] == "M")) {
				result++
			}
		}
	}

	fmt.Println("part2:", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func countAll(lines []string) int {
	r1 := regexp.MustCompile("XMAS")
	r2 := regexp.MustCompile("SAMX")
	count := 0

	for _, line := range lines {
		count += len(r1.FindAllStringIndex(line, -1)) + len(r2.FindAllStringIndex(line, -1))
	}

	return count
}

func parseData(data []byte) []string {
	result := []string{}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		result = append(result, strings.ReplaceAll(line, "\r", ""))
	}

	return result
}

func getVertical(lines []string) []string {
	result := []string{}

	for i := range len(lines[0]) {
		vertLine := ""
		for _, line := range lines {
			vertLine += string(line[i])
		}

		result = append(result, vertLine)
	}

	return result
}

func getDiagonal(lines []string) []string {
	const maxLen = 4
	// r1 := regexp.MustCompile("XMAS")
	// r2 := regexp.MustCompile("SAMX")

	result := []string{}

	for i, line := range lines {
		// left to right
		for j := range len(line) {
			diagLine := ""
			for k := 0; k < len(lines); k++ {
				row := k + i
				col := j + k

				if col >= len(line) || row >= len(lines) {
					break
				}

				diagLine += string(lines[row][col])
			}

			if len(diagLine) < maxLen {
				break
			}

			// if r1.MatchString(diagLine) || r2.MatchString(diagLine) {
			// 	fmt.Println(diagLine, "left", "i:", i, "j:", j)
			// }

			result = append(result, diagLine)

			//do full construction only on the first line
			//on the other lines, do just FIRST columns
			if i >= 1 {
				break
			}
		}

		// right to left
		for j := len(line) - 1; j >= 0; j-- {
			diagLine := ""
			for k := 0; k < len(lines); k++ {
				row := k + i
				col := j - k

				if col < 0 || row >= len(lines) {
					break
				}

				diagLine += string(lines[row][col])
			}

			if len(diagLine) < maxLen {
				break
			}

			// if r1.MatchString(diagLine) || r2.MatchString(diagLine) {
			// 	fmt.Println(diagLine, "right", "i:", i, "j:", j)
			// }

			result = append(result, diagLine)

			//do full construction only on the first line
			//on the other lines, do just LAST columns
			if i >= 1 {
				break
			}
		}
	}

	return result
}
