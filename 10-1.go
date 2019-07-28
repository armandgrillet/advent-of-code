package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inFile, err := os.Open("10.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	sliceX, sliceY, sliceVX, sliceVY := []int{}, []int{}, []int{}, []int{}

	for scanner.Scan() {
		re := regexp.MustCompile("(-)?\\d+")
		stepSlice := re.FindAllString(scanner.Text(), -1)
		strX, strY, strVX, strVY := stepSlice[0], stepSlice[1], stepSlice[2], stepSlice[3]
		x, err := strconv.Atoi(strX)
		if err != nil {
			os.Exit(1)
		}
		y, err := strconv.Atoi(strY)
		if err != nil {
			os.Exit(1)
		}
		vX, err := strconv.Atoi(strVX)
		if err != nil {
			os.Exit(1)
		}
		vY, err := strconv.Atoi(strVY)
		if err != nil {
			os.Exit(1)
		}

		sliceX = append(sliceX, x)
		sliceVX = append(sliceVX, vX)
		sliceY = append(sliceY, y)
		sliceVY = append(sliceVY, vY)
	}

	for {
		minX, minY := math.MaxInt32, math.MaxInt32
		maxX, maxY := math.MinInt32, math.MinInt32
		for i := 0; i < len(sliceX); i++ {
			sliceX[i] += sliceVX[i]
			sliceY[i] += sliceVY[i]
			switch x := sliceX[i]; {
			case x < minX:
				minX = x
			case x > maxX:
				maxX = x
			}
			switch y := sliceY[i]; {
			case y < minY:
				minY = y
			case y > maxY:
				maxY = y
			}
		}

		if math.Abs(float64(maxY-minY)) < float64(10) {
			for i := minY; i <= maxY; i++ {
				for j := minX; j <= maxX; j++ {
					symbol := " "
					for k := 0; k < len(sliceX); k++ {
						if sliceX[k] == j && sliceY[k] == i {
							symbol = "#"
						}
					}
					fmt.Print(symbol)
				}
				fmt.Println("")
			}
			os.Exit(0)
		}
	}
}
