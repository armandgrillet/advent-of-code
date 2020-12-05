package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	alphabet := make([]string, 26)
	for i := range alphabet {
		alphabet[i] = string('a' + byte(i))
	}

	var line string
	var twos, threes int
	var sawTwos, sawThrees bool

	for scanner.Scan() {
		line = scanner.Text()
		sawTwos = false
		sawThrees = false
		for _, letter := range alphabet {
			c := strings.Count(line, letter)
			if !sawTwos && c == 2 {
				twos++
				sawTwos = true
			}
			if !sawThrees && c == 3 {
				threes++
				sawThrees = true
			}
		}
	}
	fmt.Println(twos * threes)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ids := []string{}
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	idLen := len(ids[0])
	for i := 0; i < idLen; i++ {
		truncatedIDs := make(map[string]bool, len(ids))
		for _, id := range ids {
			truncatedID := id[:i] + id[i+1:]
			if truncatedIDs[truncatedID] {
				fmt.Println(truncatedID)
				return
			}
			truncatedIDs[truncatedID] = true
		}
	}
}
