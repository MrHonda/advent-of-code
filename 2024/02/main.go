package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	type_none       = iota
	type_increasing = iota
	type_decreasing = iota
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	lines := strings.Split(string(data), "\n")

	validIndexes := []int{}

	for i, line := range lines {
		values := strings.Split(string(strings.ReplaceAll(line, "\r", "")), " ")

		var prevValue int
		valuesType := type_none
		valid := true

		for j, value := range values {
			parsedValue, error := strconv.Atoi(value)
			handleError(error)

			if j == 0 {
				prevValue = parsedValue
				continue
			}
			diff := parsedValue - prevValue

			if valuesType == type_none {
				if isIncreasing(diff) {
					valuesType = type_increasing
				} else if isDecreasing(diff) {
					valuesType = type_decreasing
				} else {
					valid = false
					break
				}
			} else {
				if !(valuesType == type_increasing && isIncreasing(diff) || valuesType == type_decreasing && isDecreasing(diff)) {
					valid = false
					break
				}
			}

			prevValue = parsedValue
		}

		if valid {
			validIndexes = append(validIndexes, i)
		}
	}

	fmt.Println("valid reports:", len(validIndexes))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func isIncreasing(diff int) bool {
	return diff >= 1 && diff <= 3
}

func isDecreasing(diff int) bool {
	return diff >= -3 && diff <= -1
}
