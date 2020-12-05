package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}

	treeSlice := strings.Split(string(content), " ")
	tree := []int{}
	for _, el := range treeSlice {
		num, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tree = append(tree, num)
	}

	_, sum := partOneSumTree(0, tree)
	fmt.Println(sum)
}

func partOneSumTree(pointer int, tree []int) (int, int) {
	childNodes := tree[pointer]
	metadataEntries := tree[pointer+1]
	pointer += 2
	sum := 0
	for childNodes != 0 {
		newPointer, childSum := partOneSumTree(pointer, tree)
		pointer = newPointer
		sum += childSum
		childNodes--
	}
	for _, metadata := range tree[pointer : pointer+metadataEntries] {
		sum += metadata
	}
	return pointer + metadataEntries, sum
}

func partTwo() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}

	treeSlice := strings.Split(string(content), " ")
	tree := []int{}
	for _, el := range treeSlice {
		num, err := strconv.Atoi(el)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tree = append(tree, num)
	}

	_, sum := partTwoSumTree(0, tree)
	fmt.Println(sum)
}

func partTwoSumTree(pointer int, tree []int) (int, int) {
	childNodesCounter := tree[pointer]
	metadataEntries := tree[pointer+1]
	pointer += 2
	sum := 0

	if childNodesCounter == 0 {
		for _, metadata := range tree[pointer : pointer+metadataEntries] {
			sum += metadata
		}
		fmt.Println(sum)
		return pointer + metadataEntries, sum
	}

	childNodes := []int{}

	for childNodesCounter != 0 {
		newPointer, childSum := partTwoSumTree(pointer, tree)
		pointer = newPointer
		childNodes = append(childNodes, childSum)
		childNodesCounter--
	}

	for _, metadata := range tree[pointer : pointer+metadataEntries] {
		fmt.Println(metadata)
		if metadata <= len(childNodes) {
			sum += childNodes[metadata-1]
		}
	}
	return pointer + metadataEntries, sum
}
