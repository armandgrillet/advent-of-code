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

	freq := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			os.Exit(1)
		}
		freq += i
	}
	fmt.Println(freq)
}
