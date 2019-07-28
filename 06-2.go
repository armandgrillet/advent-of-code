package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x int
	y int
}

func main() {
	inFile, err := os.Open("06-2.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// All the elements on the corners are infinite.
	var maxX, maxY int
	minX, minY := math.MaxInt32, math.MaxInt32
	coords := make(map[coord]bool)
	for scanner.Scan() {
		re := regexp.MustCompile("[0-9]+")
		rawCoordSlice := re.FindAllString(scanner.Text(), -1)

		coordSlice := []int{}
		for _, strVal := range rawCoordSlice {
			val, err := strconv.Atoi(strVal)
			if err != nil {
				panic(err)
			}
			coordSlice = append(coordSlice, val)
		}
		c := coord{x: coordSlice[0], y: coordSlice[1]}

		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}

		coords[c] = true
	}

	smallCoords := make(map[coord]bool)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			p := coord{x: x, y: y}
			dist := 0
			for c := range coords {
				dist += int(math.Abs(float64(p.x-c.x)) + math.Abs(float64(p.y-c.y)))
			}

			if dist < 10000 {
				smallCoords[p] = true
			}
		}
	}

	fmt.Println(len(smallCoords))
	// 6537 is too high
	// 4472 is too high
	// 4284 is the right answer
}
