package main

//https://www.youtube.com/watch?v=layyhtQQuM0

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	stones, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(stones)))
}

func part1(stones []string) {
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println("part1", len(stones))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) []string {
	return strings.Split(data, " ")
}

func blink(stones []string) []string {
	for i := 0; i < len(stones); i++ {
		stone := stones[i]
		stoneInt, err := strconv.Atoi(stone)
		handleError(err)

		if stone == "0" {
			stones[i] = "1"
		} else if len(stone)%2 == 0 {
			stone1 := stone[:(len(stone) / 2)]
			stone2 := stone[(len(stone) / 2):]
			stone2 = strings.TrimLeft(stone2, "0")

			if stone2 == "" {
				stone2 = "0"
			}

			stones[i] = stone1
			stones = slices.Insert(stones, i+1, stone2)
			i++
		} else {
			stones[i] = strconv.Itoa(stoneInt * 2024)
		}
	}

	return stones
}
