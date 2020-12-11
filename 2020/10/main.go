package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var adapters []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		adapter, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	joltage := 0
	var oneDiff, threeDiff int
	for _, adapter := range adapters {
		if adapter-joltage == 1 {
			oneDiff++
		} else if adapter-joltage == 3 {
			threeDiff++
		}
		joltage = adapter
	}
	fmt.Printf("%d * %d = %d\n", oneDiff, threeDiff, oneDiff*threeDiff)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	adapters := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		adapter, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	connections := map[int]int{0: 1}
	for _, i := range adapters {
		connections[i] = connections[i-1] + connections[i-2] + connections[i-3]
		fmt.Println(i, connections[i])
	}
	fmt.Println(connections[adapters[len(adapters)-1]])
}
