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
	part2(parseData(string(stones)))
}

func part1(stones []string) {
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println("part1", len(stones))
}

func part2(stones []string) {
	stoneCounts := make(map[string]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 0; i < 75; i++ {
		stoneCounts = blink2(stoneCounts)
	}

	result := 0
	for _, count := range stoneCounts {
		result += count
	}

	fmt.Println("part2", result)
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
			stoneInt, err := strconv.Atoi(stone)
			handleError(err)
			stones[i] = strconv.Itoa(stoneInt * 2024)
		}
	}

	return stones
}

func blink2(stones map[string]int) map[string]int {
	newStones := map[string]int{}

	for stone, count := range stones {
		if stone == "0" {
			newStones["1"] += count
		} else if len(stone)%2 == 0 {
			stone1 := stone[:(len(stone) / 2)]
			stone2 := stone[(len(stone) / 2):]
			stone2 = strings.TrimLeft(stone2, "0")

			if stone2 == "" {
				stone2 = "0"
			}

			newStones[stone1] += count
			newStones[stone2] += count
		} else {
			stoneInt, err := strconv.Atoi(stone)
			handleError(err)
			newStones[strconv.Itoa(stoneInt*2024)] += count
		}
	}

	return newStones
}
