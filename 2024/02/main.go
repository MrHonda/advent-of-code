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
	part1()
	part2()
}

func part1() {
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

	fmt.Println("part 1:", len(validIndexes))
}

func part2() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	lines := strings.Split(string(data), "\n")

	validIndexes := []int{}

	for i, line := range lines {
		values := parseLine(line)

		valid := validateValues(values, true)

		if valid {
			validIndexes = append(validIndexes, i)
		} else {
			fmt.Println(line)
		}
	}

	fmt.Println("part 2:", len(validIndexes))
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

func getType(diff int) int {
	if isIncreasing(diff) {
		return type_increasing
	}
	if isDecreasing(diff) {
		return type_decreasing
	}

	return type_none
}

func isValidType(prevType int, currentType int) bool {
	return (prevType == type_none && currentType != type_none) || (prevType != type_none && prevType == currentType)
}

func parseLine(line string) []int {
	values := strings.Split(string(strings.ReplaceAll(line, "\r", "")), " ")
	result := []int{}

	for _, value := range values {
		parsedValue, error := strconv.Atoi(value)
		handleError(error)
		result = append(result, parsedValue)
	}

	return result
}

func validateValues(values []int, tryAgain bool) bool {
	prevType := type_none

	for j := range values {
		if j == 0 {
			continue
		}

		currentValue := values[j]
		prevValue := values[j-1]

		currentType := getType(currentValue - prevValue)

		if isValidType(prevType, currentType) {
			prevType = currentType
		} else if tryAgain == true {
			//brute force
			for k := range values {
				if validateValues(updateSlice(values, k), false) {
					return true
				}
			}
			return false
		} else {
			return false
		}
	}

	return true
}

func updateSlice(row []int, index int) []int {
	row2 := make([]int, len(row))
	copy(row2, row)
	slice := append(row2[:index], row2[index+1:]...)
	return slice
}
