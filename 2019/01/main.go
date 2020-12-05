package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	partTwo()
}

func partOne() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	freq := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(1)
		}
		freq += (int(math.Floor(float64(i)/3.0)) - 2)
	}
	fmt.Println(freq)
}

func partTwo() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	reqs := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(1)
		}

		freq := (int(math.Floor(float64(i)/3.0)) - 2)
		for freq > 0 {
			reqs += freq
			freq = (int(math.Floor(float64(freq)/3.0)) - 2)
		}
	}
	fmt.Println(reqs)
}
