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
		panic(err)
	}

	var seats [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, strings.Split(scanner.Text(), ""))
	}

	var unchanged, newUnchanged int
	for {
		newUnchanged = 0
		newSeats := make([][]string, len(seats))
		for i := 0; i < len(seats); i++ {
			for j := 0; j < len(seats[i]); j++ {
				if seats[i][j] == "L" && adjacentOccupiedSeats(i, j, seats) == 0 {
					newSeats[i] = append(newSeats[i], "#")
					newUnchanged++
				} else if seats[i][j] == "#" && adjacentOccupiedSeats(i, j, seats) >= 4 {
					newSeats[i] = append(newSeats[i], "L")
					newUnchanged++
				} else {
					newSeats[i] = append(newSeats[i], seats[i][j])
				}
			}
		}

		if newUnchanged == unchanged {
			occupied := 0
			for _, row := range newSeats {
				for _, seat := range row {
					if seat == "#" {
						occupied++
					}
				}
			}
			fmt.Println(occupied)
			return
		}
		unchanged = newUnchanged
		seats = newSeats
	}
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var seats [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seats = append(seats, strings.Split(scanner.Text(), ""))
	}

	var unchanged, newUnchanged int
	for {
		newUnchanged = 0
		newSeats := make([][]string, len(seats))
		for i := 0; i < len(seats); i++ {
			for j := 0; j < len(seats[i]); j++ {
				if seats[i][j] == "L" && directionalOccupiedSeats(i, j, seats) == 0 {
					newSeats[i] = append(newSeats[i], "#")
					newUnchanged++
				} else if seats[i][j] == "#" && directionalOccupiedSeats(i, j, seats) >= 5 {
					newSeats[i] = append(newSeats[i], "L")
					newUnchanged++
				} else {
					newSeats[i] = append(newSeats[i], seats[i][j])
				}
			}
		}

		if newUnchanged == unchanged {
			occupied := 0
			for _, row := range newSeats {
				for _, seat := range row {
					if seat == "#" {
						occupied++
					}
				}
			}
			fmt.Println(occupied)
			return
		}
		unchanged = newUnchanged
		seats = newSeats
	}
}

func adjacentOccupiedSeats(row int, col int, seats [][]string) int {
	count := 0
	if row-1 >= 0 && col-1 >= 0 && seats[row-1][col-1] == "#" {
		count++
	}
	if row-1 >= 0 && seats[row-1][col] == "#" {
		count++
	}
	if row-1 >= 0 && col+1 < len(seats[row]) && seats[row-1][col+1] == "#" {
		count++
	}

	if col-1 >= 0 && seats[row][col-1] == "#" {
		count++
	}
	if col+1 < len(seats[row]) && seats[row][col+1] == "#" {
		count++
	}

	if row+1 < len(seats) && col-1 >= 0 && seats[row+1][col-1] == "#" {
		count++
	}
	if row+1 < len(seats) && seats[row+1][col] == "#" {
		count++
	}
	if row+1 < len(seats) && col+1 < len(seats[row+1]) && seats[row+1][col+1] == "#" {
		count++
	}

	return count
}

func directionalOccupiedSeats(row int, col int, seats [][]string) int {
	count := 0
	var seat string

	for down := row - 1; down >= 0; down-- {
		seat = seats[down][col]
		if seat == "#" {
			count++
			break
		} else if seat == "L" {
			break
		}
	}
	for right := col + 1; right < len(seats[row]); right++ {
		seat = seats[row][right]
		if seat == "#" {
			count++
			break
		} else if seat == "L" {
			break
		}
	}
	for up := row + 1; up < len(seats); up++ {
		seat = seats[up][col]
		if seat == "#" {
			count++
			break
		} else if seat == "L" {
			break
		}
	}
	for left := col - 1; left >= 0; left-- {
		seat = seats[row][left]
		if seat == "#" {
			count++
			break
		} else if seat == "L" {
			break
		}
	}

	for down := row - 1; down >= 0; down-- {
		right := col + (row - down)
		if right < len(seats[row]) {
			seat = seats[down][right]
			if seat == "#" {
				count++
				break
			} else if seat == "L" {
				break
			}
		}
	}

	for up := row + 1; up < len(seats); up++ {
		right := col + (up - row)
		if right < len(seats[row]) {
			seat = seats[up][right]
			if seat == "#" {
				count++
				break
			} else if seat == "L" {
				break
			}
		}
	}

	for up := row + 1; up < len(seats); up++ {
		left := col - (up - row)
		if left >= 0 {
			seat = seats[up][left]
			if seat == "#" {
				count++
				break
			} else if seat == "L" {
				break
			}
		}
	}

	for down := row - 1; down >= 0; down-- {
		left := col - (row - down)
		if left >= 0 {
			seat = seats[down][left]
			if seat == "#" {
				count++
				break
			} else if seat == "L" {
				break
			}
		}
	}

	return count
}
