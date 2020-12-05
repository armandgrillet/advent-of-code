package main

import (
	"fmt"
)

func main() {
	players := make(map[int]int)
	for i := 0; i <= 405; i++ {
		players[i] = 0
	}

	marbles := []int{0, 2, 1}
	currentMarble := 1 // marbles[1] == 2

	for i := 3; i <= 7095300; i++ {
		if i%100000 == 0 {
			fmt.Println("1/70")
		}
		if i%23 != 0 {
			newMarble := (currentMarble + 2)
			if newMarble > len(marbles) {
				newMarble = newMarble % len(marbles)
			}
			if newMarble < len(marbles) {
				marbles = append(marbles, 0)
				copy(marbles[newMarble+1:], marbles[newMarble:])
				marbles[newMarble] = i
			} else if newMarble == len(marbles) {
				marbles = append(marbles, i)
			}
			currentMarble = newMarble
		} else {
			currentPlayer := i % (len(players) - 1)
			players[currentPlayer] += i
			marbleToTake := (currentMarble - 7)
			if marbleToTake < 0 {
				marbleToTake = len(marbles) + marbleToTake
			}
			players[currentPlayer] += marbles[marbleToTake]
			copy(marbles[marbleToTake:], marbles[marbleToTake+1:])
			marbles[len(marbles)-1] = 0
			marbles = marbles[:len(marbles)-1]
			currentMarble = marbleToTake
			if currentMarble == len(marbles) {
				currentMarble = 0
			}
		}
	}
	maxScore := 0
	for _, score := range players {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}
