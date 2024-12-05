package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	type_none   = iota
	type_before = iota
	type_after  = iota
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
	part2(parseData(string(data)))
}

func part1(rules [][]int, orders [][]int) {
	result := 0

	for _, order := range orders {
		valid := true

		for i, val1 := range order {
			for j, val2 := range order {
				if i == j {
					continue
				}

				if (j < i && !canBeBefore(val1, val2, rules)) || (j > i && !canBeAfter(val1, val2, rules)) {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			middle := order[(len(order)-1)/2]
			// fmt.Println(order, "middle", middle)
			result += middle
		}
	}

	fmt.Println("part1", result)
}

func part2(rules [][]int, orders [][]int) {
	result := 0

	for _, order := range orders {
		valid := true
		counter := 0

		for {
			validOrder, index1, index2, invalidType := validateOrder(order, rules)

			if validOrder {
				break
			}

			valid = false

			if invalidType == type_after {
				tmp := order[index2]
				order = slices.Delete(order, index2, index2+1)
				order = slices.Insert(order, index1, tmp)
			} else if invalidType == type_before {
				tmp := order[index1]
				order = slices.Delete(order, index1, index1+1)
				order = slices.Insert(order, index2, tmp)
			} else {
				panic("invalid invalidType")
			}

			counter++

			// fmt.Println(counter)
		}

		if !valid {
			middle := order[(len(order)-1)/2]
			// fmt.Println(order, "middle", middle)
			result += middle
		}
	}

	fmt.Println("part2", result)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) ([][]int, [][]int) {
	rules := [][]int{}
	orders := [][]int{}

	splitted := strings.Split(data, "#")

	rulesLinesRaw := strings.Split(splitted[0], "\r\n")

	for _, ruleLine := range rulesLinesRaw {
		if ruleLine == "" {
			continue
		}

		rowRules := strings.Split(ruleLine, "|")

		val1, err := strconv.Atoi(rowRules[0])
		handleError(err)
		val2, err := strconv.Atoi(rowRules[1])
		handleError(err)
		rules = append(rules, []int{val1, val2})
	}

	ordersLinesRaw := strings.Split(splitted[1], "\r\n")

	for _, orderLine := range ordersLinesRaw {
		if orderLine == "" {
			continue
		}

		rowOrders := strings.Split(orderLine, ",")

		ordersRow := []int{}

		for _, rowOrder := range rowOrders {
			val, err := strconv.Atoi(rowOrder)
			handleError(err)
			ordersRow = append(ordersRow, val)
		}

		orders = append(orders, ordersRow)
	}

	return rules, orders
}

func canBeAfter(val1 int, val2 int, rules [][]int) bool {
	for _, rule := range rules {
		if rule[0] == val1 && rule[1] == val2 {
			return true
		}

		if rule[0] == val2 && rule[1] == val1 {
			return false
		}
	}

	return true
}

func canBeBefore(val1 int, val2 int, rules [][]int) bool {
	for _, rule := range rules {
		if rule[0] == val1 && rule[1] == val2 {
			return false
		}

		if rule[0] == val2 && rule[1] == val1 {
			return true
		}
	}

	return true
}

// valid, index1, index2, type
func validateOrder(order []int, rules [][]int) (bool, int, int, int) {
	for i, val1 := range order {
		for j, val2 := range order {
			if i == j {
				continue
			}

			if j < i && !canBeBefore(val1, val2, rules) {
				return false, i, j, type_before
			}

			if j > i && !canBeAfter(val1, val2, rules) {
				return false, i, j, type_after
			}
		}
	}

	return true, -1, -1, type_none
}
