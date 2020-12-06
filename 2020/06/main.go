package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partOne()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var line string

	answers := make(map[string]bool)
	sum := 0

	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			for _, answer := range line {
				answers[string(answer)] = true
			}
		} else {
			sum += len(answers)
			answers = make(map[string]bool)
		}
	}
	fmt.Println(sum)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var line string

	answers := make(map[string]int, 26)
	groupSize := 0
	sum := 0

	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			groupSize++
			for _, answer := range line {
				answers[string(answer)]++
			}
		} else {
			for _, answeredYes := range answers {
				if answeredYes == groupSize {
					sum++
				}
			}

			answers = make(map[string]int, 26)
			groupSize = 0
		}
	}

	fmt.Println(sum)
}
