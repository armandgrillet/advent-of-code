package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	direction := "E"
	positions := map[string]int{
		"E": 0,
		"S": 0,
		"W": 0,
		"N": 0,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var stepDirection string
	var stepDistanceOrDegrees int
	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = strings.SplitN(scanner.Text(), "", 2)

		stepDistanceOrDegrees, err = strconv.Atoi(instructions[1])
		if err != nil {
			panic(err)
		}

		if instructions[0] != "L" && instructions[0] != "R" {
			stepDirection = direction
			if instructions[0] == "N" || instructions[0] == "S" || instructions[0] == "E" || instructions[0] == "W" {
				stepDirection = instructions[0]
			}
			positions[stepDirection] += stepDistanceOrDegrees
		} else {
			direction = newDirection(direction, instructions[0], stepDistanceOrDegrees)
		}
	}

	fmt.Println(abs(positions["E"]-positions["W"]) + abs(positions["S"]-positions["N"]))
}

func partTwo() {
	shipPositions := map[string]int{
		"E": 0,
		"N": 0,
	}
	waypointPositions := map[string]int{
		"E": 10,
		"N": 1,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var stepDistanceOrDegrees int
	var instructions []string
	scanner := bufio.NewScanner(file)

	var distE, distN int
	for scanner.Scan() {
		instructions = strings.SplitN(scanner.Text(), "", 2)

		stepDistanceOrDegrees, err = strconv.Atoi(instructions[1])
		if err != nil {
			panic(err)
		}

		if instructions[0] == "N" || instructions[0] == "E" {
			waypointPositions[instructions[0]] += stepDistanceOrDegrees
		} else if instructions[0] == "S" {
			waypointPositions["N"] -= stepDistanceOrDegrees
		} else if instructions[0] == "W" {
			waypointPositions["E"] -= stepDistanceOrDegrees
		} else {
			distE = waypointPositions["E"] - shipPositions["E"]
			distN = waypointPositions["N"] - shipPositions["N"]

			if instructions[0] == "F" {
				waypointPositions["E"] += stepDistanceOrDegrees * distE
				waypointPositions["N"] += stepDistanceOrDegrees * distN
				shipPositions["E"] = waypointPositions["E"] - distE
				shipPositions["N"] = waypointPositions["N"] - distN
			} else if instructions[0] == "L" || instructions[0] == "R" {
				if (instructions[0] == "R" && stepDistanceOrDegrees == 90) || (instructions[0] == "L" && stepDistanceOrDegrees == 270) {
					waypointPositions["E"] = shipPositions["E"] + distN
					waypointPositions["N"] = shipPositions["N"] - distE
				} else if (instructions[0] == "R" && stepDistanceOrDegrees == 270) || (instructions[0] == "L" && stepDistanceOrDegrees == 90) {
					waypointPositions["E"] = shipPositions["E"] - distN
					waypointPositions["N"] = shipPositions["N"] + distE
				} else if stepDistanceOrDegrees == 180 {
					waypointPositions["E"] = shipPositions["E"] - distE
					waypointPositions["N"] = shipPositions["N"] - distN
				}
			}
		}
	}
	fmt.Println(abs(shipPositions["E"]) + abs(shipPositions["N"]))
}

func newDirection(direction string, side string, degrees int) string {
	var degreesToSlice int
	if side == "L" {
		switch degrees {
		case 90:
			degreesToSlice = 3
		case 270:
			degreesToSlice = 1
		default:
			degreesToSlice = degrees / 90
		}
	} else {
		degreesToSlice = degrees / 90
	}

	rightDirections := map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}

	newDirection := (rightDirections[direction] + degreesToSlice) % 4

	for k, v := range rightDirections {
		if v == newDirection {
			return k
		}
	}
	return ""
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
