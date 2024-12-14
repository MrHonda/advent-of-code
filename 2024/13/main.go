package main

//https://www.youtube.com/watch?v=-5J-DAsWuJc

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(blocks [][6]int) {
	result := 0

	for _, block := range blocks {
		ax, ay, bx, by, px, py := block[0], block[1], block[2], block[3], block[4], block[5]
		ca := float64(px*by-py*bx) / float64(ax*by-ay*bx)
		cb := (float64(px) - float64(ax)*ca) / float64(bx)

		if isInt(ca) && isInt(cb) {
			result += int(ca)*3 + int(cb)*1
		}
	}

	fmt.Println("part1", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][6]int {
	result := [][6]int{}
	blocks := strings.Split(data, "\r\n\r\n")

	re := regexp.MustCompile(`\d+`)

	for _, block := range blocks {
		matches := re.FindAllString(block, -1)

		// Convert to integers
		var numbers [6]int
		for i, match := range matches {
			num, _ := strconv.Atoi(match) // Convert string to integer
			numbers[i] = num
		}

		result = append(result, numbers)
	}

	return result
}

func isInt(f float64) bool {
	return f == float64(int(f))
}
