package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("02.txt")
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
