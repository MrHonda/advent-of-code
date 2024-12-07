package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Calculation struct {
	result  int
	numbers []int
}

const (
	op_add      = "+"
	op_multiple = "*"
)

var operators = []string{op_add, op_multiple}

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(calculations []Calculation) {
	result := 0

	for _, calculation := range calculations {
		combinations := getCombinations(len(calculation.numbers) - 1)

		for _, combination := range combinations {
			val := 0
			for i, op := range combination {
				if i == 0 {
					val = calculate(calculation.numbers[i], calculation.numbers[i+1], op)
				} else {
					val = calculate(val, calculation.numbers[i+1], op)
				}

				if val > calculation.result {
					break
				}
			}

			if val == calculation.result {
				result += val
				break
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

func parseData(data string) []Calculation {
	lines := strings.Split(string(data), "\r\n")
	calculations := []Calculation{}

	for _, line := range lines {
		split := strings.Split(line, ":")
		result, err := strconv.Atoi(split[0])
		handleError(err)
		numbersLine := strings.TrimSpace(split[1])
		numbersString := strings.Split(numbersLine, " ")

		numbers := []int{}

		for _, rawNumber := range numbersString {
			val, err := strconv.Atoi(rawNumber)
			handleError(err)
			numbers = append(numbers, val)
		}

		calculations = append(calculations, Calculation{result: result, numbers: numbers})
	}

	return calculations
}

func calculate(val1 int, val2 int, operator string) int {
	switch operator {
	case op_add:
		return val1 + val2
	case op_multiple:
		return val1 * val2
	default:
		panic("invalid operator" + operator)
	}
}

func getCombinations(count int) [][]string {
	totalCombinations := int(math.Pow(float64(len(operators)), float64(count)))
	combinations := make([][]string, 0, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combo := make([]string, count)
		temp := i
		for j := 0; j < count; j++ {
			combo[j] = operators[temp%len(operators)]
			temp /= len(operators)
		}
		combinations = append(combinations, combo)
	}

	return combinations
}
