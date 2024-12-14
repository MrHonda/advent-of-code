package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const width = 101
const height = 103
const middle_x = width / 2
const middle_y = height / 2

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(robots [][4]int) {
	quadrants := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		pos := move(robot, 100)

		if pos[0] < middle_x && pos[1] < middle_y { //left top
			quadrants[0]++
		} else if pos[0] > middle_x && pos[1] < middle_y { // right top
			quadrants[1]++
		} else if pos[0] < middle_x && pos[1] > middle_y { // left bottom
			quadrants[2]++
		} else if pos[0] > middle_x && pos[1] > middle_y { // right bottom
			quadrants[3]++
		}
	}

	result := 1

	for _, quadrant := range quadrants {
		result *= quadrant
	}

	fmt.Println("part1", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][4]int {
	result := [][4]int{}
	robots := strings.Split(data, "\r\n")

	re := regexp.MustCompile(`-?\d+`)

	for _, robot := range robots {
		matches := re.FindAllString(robot, -1)
		// Convert to integers
		var numbers [4]int
		for i, match := range matches {
			num, _ := strconv.Atoi(match) // Convert string to integer
			numbers[i] = num
		}

		result = append(result, numbers)
	}

	return result
}

func move(robot [4]int, times int) [2]int {
	pos := [2]int{robot[0], robot[1]}
	dir := [2]int{robot[2], robot[3]}

	for range times {
		x := (pos[0] + dir[0] + width) % width
		y := (pos[1] + dir[1] + height) % height

		pos[0] = x
		pos[1] = y

		// fmt.Println(x, y)
	}

	return pos
}
