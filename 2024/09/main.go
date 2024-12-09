package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	id         int
	startIndex int
	endIndex   int
}

func (file File) size() int {
	return file.endIndex - file.startIndex + 1
}

const freeSpaceVal = -1

func main() {
	data, err := os.ReadFile("input.txt")
	handleError(err)

	part1(parseData(string(data)))
	blocks, _ := parseData(string(data))
	part2(blocks)
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

func part2(blocks []int) {
	processedFiles := map[int]bool{}

	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]

		if block == freeSpaceVal {
			continue
		}

		file := findFile(blocks, i)
		fileSize := file.size()
		i -= fileSize - 1

		if processedFiles[file.id] {
			continue
		}

		if file.startIndex == 0 {
			break
		}

		processedFiles[file.id] = true

		freeStart, freeEnd := findFreeSpace(blocks, fileSize, i)

		if freeStart >= 0 && freeEnd >= 0 {
			for j := freeStart; j <= freeEnd; j++ {
				blocks[j] = file.id
			}

			for j := file.startIndex; j <= file.endIndex; j++ {
				blocks[j] = freeSpaceVal
			}
		}
	}

	fmt.Println("part2", calculateCheckSum(blocks))
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

func findFile(blocks []int, endI int) File {
	foundId := -1
	startIndex := 0
	endIndex := -1

	for i := endI; i >= 0; i-- {
		block := blocks[i]

		if block == freeSpaceVal {
			if foundId != -1 {
				startIndex = i + 1
				break
			}
		}

		if foundId == -1 {
			foundId = block
			endIndex = i
		} else if block != foundId {
			startIndex = i + 1
			break
		}
	}

	return File{id: foundId, startIndex: startIndex, endIndex: endIndex}
}

func findFreeSpace(blocks []int, size int, maxI int) (int, int) {
	startIndex := -1

	for i, block := range blocks {
		if i >= maxI {
			return -1, -1
		}
		if block != freeSpaceVal {
			startIndex = -1
			continue
		}

		if startIndex == -1 {
			startIndex = i
		}

		if i-startIndex+1 >= size {
			return startIndex, i
		}
	}

	return -1, -1
}
