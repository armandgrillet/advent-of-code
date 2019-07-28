package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("08.txt")
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

	_, sum := sumTree(0, tree)
}

func sumTree(pointer int, tree []int) (int, int) {
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
		newPointer, childSum := sumTree(pointer, tree)
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
