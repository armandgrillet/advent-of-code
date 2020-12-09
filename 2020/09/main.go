package main

import (
	"bufio"
	"fmt"
	"os"
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

	preamble := 25
	head := 25
	var numbers []int

	pos := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if pos < preamble {
			numbers = append(numbers, number)
			pos++
		} else if isSum(number, numbers) {
			if head < len(numbers) {
				numbers = append(numbers[:1], numbers[2:]...)
			}
			numbers = append(numbers, number)
		} else {
			fmt.Println(number)
			return
		}
	}
}

func isSum(number int, numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			} else if number == numbers[i]+numbers[j] {
				return true
			}
		}
	}
	return false
}

func partTwo() {
	invalidNumber := 22406676

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	var sum, smallest, biggest int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j <= len(numbers); j++ {
			sum = 0
			smallest = invalidNumber
			biggest = 0
			for _, number := range numbers[i:j] {
				if number < smallest {
					smallest = number
				}
				if number > biggest {
					biggest = number
				}

				sum += number
				if sum == invalidNumber {
					fmt.Println(smallest + biggest)
					return
				}
			}
		}
	}
}
