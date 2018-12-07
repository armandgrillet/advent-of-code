package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inFile, err := os.Open("01.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	changes := []int{}
	for scanner.Scan() {
		change, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(1)
		}

		changes = append(changes, change)
	}

	freqs := make(map[int]bool)
	freq := 0
	for {
		for _, change := range changes {
			freq += change
			if freqs[freq] {
				fmt.Println(freq)
				return
			}
			freqs[freq] = true
		}
	}
}
