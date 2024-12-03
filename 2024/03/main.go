package main

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

	part1(string(data))
}

func part1(data string) {
	r := regexp.MustCompile(`mul(\(\d+,\d+\))`)
	matches := r.FindAllStringSubmatch(data, -1)
	// fmt.Println(matches)

	result := 0

	for _, match := range matches {
		result += calculate(match[1])
	}

	fmt.Println("part1:", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func calculate(match string) int {
	match = strings.ReplaceAll(match, "(", "")
	match = strings.ReplaceAll(match, ")", "")
	values := strings.Split(match, ",")

	val1, _ := strconv.Atoi(values[0])
	val2, _ := strconv.Atoi(values[1])

	return val1 * val2
}
