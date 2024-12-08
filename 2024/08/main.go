package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
	part2(parseData(string(data)))
}

func part1(data [][]string) {
	antinodes := map[string]bool{}

	for i, line := range data {
		for j, char := range line {
			if char == "." {
				continue
			}

			findAntinodes(char, i, j, data, antinodes)
		}
	}

	// fmt.Println(antinodes)
	fmt.Println("part1", len(antinodes))
}

func part2(data [][]string) {
	antinodes := map[string]bool{}
	antennas := map[string]int{}

	// find out the number of each antenna
	for _, line := range data {
		for _, char := range line {
			if char == "." {
				continue
			}

			antennas[char]++
		}
	}

	for i, line := range data {
		for j, char := range line {
			if char == "." {
				continue
			}

			findAntinodes2(char, i, j, data, antinodes)

			// count an antenna as an antinode if there are is at least 2 of them
			if antennas[char] > 1 {
				antinodePos := getPos(i, j)

				if !antinodes[antinodePos] {
					antinodes[antinodePos] = true
				}
			}

		}
	}

	// fmt.Println(antinodes)
	fmt.Println("part2", len(antinodes))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) [][]string {
	lines := strings.Split(string(data), "\r\n")
	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func findAntinodes(antenna string, antennaI int, antennaJ int, data [][]string, antinodes map[string]bool) {
	for i := antennaI + 1; i < len(data); i++ {
		for j, char := range data[i] {
			if char != antenna {
				continue
			}

			dirI := i - antennaI
			dirJ := j - antennaJ

			antinodeI := antennaI - dirI
			antinodeJ := antennaJ - dirJ
			antinodePos := getPos(antinodeI, antinodeJ)

			if isValidPos(antinodeI, antinodeJ, data) && !antinodes[antinodePos] {
				antinodes[antinodePos] = true
			}

			antinodeI = i + dirI
			antinodeJ = j + dirJ
			antinodePos = getPos(antinodeI, antinodeJ)

			if isValidPos(antinodeI, antinodeJ, data) && !antinodes[antinodePos] {
				antinodes[antinodePos] = true
			}
		}
	}
}

func findAntinodes2(antenna string, antennaI int, antennaJ int, data [][]string, antinodes map[string]bool) {
	for i := antennaI + 1; i < len(data); i++ {
		for j, char := range data[i] {
			if char != antenna {
				continue
			}

			dirI := i - antennaI
			dirJ := j - antennaJ

			antinodeI := antennaI
			antinodeJ := antennaJ

			for {
				antinodeI -= dirI
				antinodeJ -= dirJ
				antinodePos := getPos(antinodeI, antinodeJ)

				if !isValidPos(antinodeI, antinodeJ, data) {
					break
				}

				if !antinodes[antinodePos] {
					antinodes[antinodePos] = true
				}
			}

			antinodeI = i
			antinodeJ = j

			for {
				antinodeI += dirI
				antinodeJ += dirJ
				antinodePos := getPos(antinodeI, antinodeJ)

				if !isValidPos(antinodeI, antinodeJ, data) {
					break
				}

				if !antinodes[antinodePos] {
					antinodes[antinodePos] = true
				}
			}
		}
	}
}

func getPos(i int, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func isValidPos(i int, j int, data [][]string) bool {
	return i >= 0 && j >= 0 && i < len(data) && j < len(data[0])
}
