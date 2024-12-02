package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var list1 = []int{}
var list2 = []int{}

func main() {
	part1()
	part2()
}

func part1() {
	prepareInput()

	var distance = 0

	for len(list1) > 0 {
		var num1 = find(&list1)
		var num2 = find(&list2)
		distance += absInt(num1 - num2)
	}

	println("part1:", distance)
}

func part2() {
	prepareInput()

	var distance = 0

	for _, value := range list1 {
		println(value, findAppears(list2, value))
		distance += value * findAppears(list2, value)
	}

	println("part2:", distance)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func prepareInput() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	list1 = nil
	list2 = nil

	// Split the content into lines
	lines := strings.Split(string(data), "\n")

	// Print each line
	for _, line := range lines {
		values := strings.Split(string(strings.ReplaceAll(line, "\r", "")), "   ")

		value1, err1 := strconv.Atoi(values[0])
		handleError(err1)

		value2, err2 := strconv.Atoi(values[1])
		handleError(err2)

		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}
}

func find(slice *[]int) int {
	var result = math.MaxInt
	var index = -1

	for i, v := range *slice {
		if v < result {
			result = v
			index = i
		}
	}

	*slice = append((*slice)[:index], (*slice)[index+1:]...)

	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findAppears(list []int, search int) int {
	result := 0

	for _, value := range list {
		if value == search {
			result++
		}
	}

	return result
}
