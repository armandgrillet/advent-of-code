package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inFile, err := os.Open("05.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var polymer string
	for scanner.Scan() {
		polymer = scanner.Text()
		if err != nil {
			os.Exit(1)
		}
	}

	i := 0
	for i < len(polymer)-1 {
		if polymer[i]+32 == polymer[i+1] || polymer[i]-32 == polymer[i+1] {
			polymer = polymer[:i] + polymer[i+2:]
			i = 0
		} else {
			i++
		}
	}

	fmt.Println(len(polymer))
}
