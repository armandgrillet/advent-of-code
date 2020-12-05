package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	position := 0
	treesEncountered := 0

	scanner := bufio.NewScanner(file)
	// We do not count on the first line as we haven't gone down yet.
	scanner.Scan()
	mapWidth := len(scanner.Text())
	position += 3

	for scanner.Scan() {
		if string(scanner.Text()[(position%mapWidth)]) == "#" {
			treesEncountered++
		}
		position += 3
	}

	fmt.Println(treesEncountered)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}

	fmt.Println(slopes(trees, 1, 1) * slopes(trees, 3, 1) * slopes(trees, 5, 1) * slopes(trees, 7, 1) * slopes(trees, 1, 2))
}

func slopes(trees []string, right int, down int) int {
	position := 0
	treesEncountered := 0

	if len(trees) <= down {
		return treesEncountered
	}

	mapWidth := len(trees[0])
	position += right

	for i := down; i < len(trees); i += down {
		if string(trees[i][position%mapWidth]) == "#" {
			treesEncountered++
		}
		position += right
	}
	return treesEncountered
}
