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
	partOne()
}

func partOne() {
	inFile, err := os.Open("06-1.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// All the elements on the corners are infinite.
	var maxX, maxY int
	minX, minY := math.MaxInt32, math.MaxInt32
	coords := make(map[coord]int)
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

		coords[c] = 0
	}

	// Will start at minX - 50 and go up to maxX+50
	infiniteCoords := make(map[coord]bool)
	veryMinX, veryMinY := minX-50, minY-50
	veryMaxX, veryMaxY := maxX+50, maxY+50
	for veryLowX := veryMinX; veryLowX < veryMaxX; veryLowX++ {
		pMin := coord{x: veryLowX, y: veryMinY}
		pMax := coord{x: veryLowX, y: veryMaxY}
		var matchingCoords []coord
		minDist := math.MaxInt32
		for c := range coords {
			dist := int(math.Abs(float64(pMin.x-c.x)) + math.Abs(float64(pMin.y-c.y)))
			if dist == minDist {
				matchingCoords = append(matchingCoords, c)
			} else if dist < minDist {
				matchingCoords = []coord{c}
				minDist = dist
			}
		}

		if len(matchingCoords) == 1 {
			infiniteCoords[matchingCoords[0]] = true
		}

		matchingCoords = []coord{}
		minDist = math.MaxInt32
		for c := range coords {
			dist := int(math.Abs(float64(pMax.x-c.x)) + math.Abs(float64(pMax.y-c.y)))
			if dist == minDist {
				matchingCoords = append(matchingCoords, c)
			} else if dist < minDist {
				matchingCoords = []coord{c}
				minDist = dist
			}
		}

		if len(matchingCoords) == 1 {
			infiniteCoords[matchingCoords[0]] = true
		}
	}

	for veryLowY := veryMinY; veryLowY < veryMaxY; veryLowY++ {
		pMin := coord{x: veryMinX, y: veryLowY}
		pMax := coord{x: veryMaxX, y: veryLowY}
		var matchingCoords []coord
		minDist := math.MaxInt32
		for c := range coords {
			dist := int(math.Abs(float64(pMin.x-c.x)) + math.Abs(float64(pMin.y-c.y)))
			if dist == minDist {
				matchingCoords = append(matchingCoords, c)
			} else if dist < minDist {
				matchingCoords = []coord{c}
				minDist = dist
			}
		}

		if len(matchingCoords) == 1 {
			infiniteCoords[matchingCoords[0]] = true
		}

		matchingCoords = []coord{}
		minDist = math.MaxInt32
		for c := range coords {
			dist := int(math.Abs(float64(pMax.x-c.x)) + math.Abs(float64(pMax.y-c.y)))
			if dist == minDist {
				matchingCoords = append(matchingCoords, c)
			} else if dist < minDist {
				matchingCoords = []coord{c}
				minDist = dist
			}
		}

		if len(matchingCoords) == 1 {
			infiniteCoords[matchingCoords[0]] = true
		}
	}

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			p := coord{x: x, y: y}

			if _, ok := coords[p]; ok {
				coords[p]++
				continue
			}

			var matchingCoords []coord
			minDist := math.MaxInt32
			for c := range coords {
				dist := int(math.Abs(float64(p.x-c.x)) + math.Abs(float64(p.y-c.y)))
				if dist == minDist {
					matchingCoords = append(matchingCoords, c)
				} else if dist < minDist {
					matchingCoords = []coord{c}
					minDist = dist
				}
			}

			if len(matchingCoords) >= 1 {
				if len(matchingCoords) == 1 {
					coords[matchingCoords[0]]++
				}
			}
		}
	}

	fmt.Println(infiniteCoords)

	maxTiles := 0
	for k, v := range coords {
		if _, ok := infiniteCoords[k]; !ok {
			if v > maxTiles {
				maxTiles = v
			}
		}
	}
	fmt.Println(maxTiles)
}

func partTwo() {
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
}
