package main

import "fmt"

func main() {
	partOne()
}

func partOne() {
	serial := 7803
	cells := make([][]int, 300)
	for i := range cells {
		cells[i] = make([]int, 300)
	}

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackID := x + 10
			power := rackID * y
			power += serial
			power *= rackID
			hundreds := power % 1000
			power = (hundreds - (hundreds % 100)) / 100
			power -= 5
			cells[x-1][y-1] = power
		}
	}

	max := 0

	for x := 0; x <= 297; x++ {
		for y := 0; y <= 297; y++ {
			sum := cells[x][y]
			sum += cells[x][y+1]
			sum += cells[x][y+2]
			sum += cells[x+1][y]
			sum += cells[x+1][y+1]
			sum += cells[x+1][y+2]
			sum += cells[x+2][y]
			sum += cells[x+2][y+1]
			sum += cells[x+2][y+2]

			if sum > max {
				fmt.Printf("%d,%d\n", x+1, y+1)
				max = sum
			}
		}
	}

}
