package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partTwo()
}

func seatFinder(min int, max int, minMatch string, maxMatch string, chain string) int {
	if min == max {
		return min
	}

	dist := (max - min) + 1
	if dist%2 != 0 {
		panic("Distance is not divisible by 2")
	}

	if chain == "" {
		panic("The chain cannot be parsed")
	}

	switch string(chain[0]) {
	case minMatch: // F or L
		return seatFinder(min, max-(dist/2), minMatch, maxMatch, chain[1:])
	case maxMatch: // B or R
		return seatFinder(min+(dist/2), max, minMatch, maxMatch, chain[1:])
	default:
		return seatFinder(min, max, minMatch, maxMatch, chain[1:])
	}
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	highestSeat := 0
	maxRow := 127
	maxCol := 7
	var row, col int

	for scanner.Scan() {
		row = seatFinder(0, maxRow, "F", "B", scanner.Text())
		col = seatFinder(0, maxCol, "L", "R", scanner.Text())
		if row*(maxCol+1)+col > highestSeat {
			highestSeat = row*(maxCol+1) + col
		}
	}
	fmt.Println(highestSeat)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	maxRow := 127
	maxCol := 7
	var row, col int

	seats := make([][]bool, maxRow+1)
	for i := 0; i < maxRow+1; i++ {
		seats[i] = make([]bool, maxCol+1)
	}

	for scanner.Scan() {
		row = seatFinder(0, maxRow, "F", "B", scanner.Text())
		col = seatFinder(0, maxCol, "L", "R", scanner.Text())
		seats[row][col] = true
	}

	for i := 0; i < maxRow+1; i++ {
		for j := 1; j < maxCol; j++ {
			if seats[i][j-1] == true && seats[i][j] == false && seats[i][j+1] == true {
				fmt.Println(i*(maxCol+1) + j)
			}
		}
	}
}
