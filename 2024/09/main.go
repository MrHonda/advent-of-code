package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const freeSpaceVal = -1

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
}

func part1(blocks []int, freeBlocks int) {
	for freeBlocks > 0 {
		block := blocks[len(blocks)-1]
		blocks = blocks[:len(blocks)-1]
		freeBlocks--

		if block == freeSpaceVal {
			continue
		}

		freeBlockIndex := findFreeBlockIndex(blocks)
		blocks[freeBlockIndex] = block
	}

	fmt.Println("part1", calculateCheckSum(blocks))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseData(data string) ([]int, int) {
	digits := strings.Split(string(data), "")
	blocks := []int{}

	isFile := true
	fileIndex := 0
	totalFreeBlocks := 0

	for _, digit := range digits {
		val, err := strconv.Atoi(digit)
		handleError(err)

		for range val {
			if isFile {
				blocks = append(blocks, fileIndex)
			} else {
				blocks = append(blocks, freeSpaceVal)
				totalFreeBlocks++
			}
		}

		if isFile {
			isFile = false
		} else {
			isFile = true
			fileIndex++
		}
	}

	return blocks, totalFreeBlocks
}

func findFreeBlockIndex(blocks []int) int {
	for i, block := range blocks {
		if block == freeSpaceVal {
			return i
		}
	}

	panic("no free block found")
}

func calculateCheckSum(blocks []int) int {
	result := 0
	for i, block := range blocks {
		if block == freeSpaceVal {
			continue
		}

		result += i * block
	}

	return result
}
