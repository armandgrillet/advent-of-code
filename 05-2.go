package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	shortestPolymer := len(polymer)
	for letter := 0; letter < 26; letter++ {
		improvedPolymer := strings.Replace(polymer, string('A'+letter), "", -1)
		improvedPolymer = strings.Replace(improvedPolymer, string('a'+letter), "", -1)
		i := 0
		for i < len(improvedPolymer)-1 {
			if improvedPolymer[i]+32 == improvedPolymer[i+1] || improvedPolymer[i]-32 == improvedPolymer[i+1] {
				improvedPolymer = improvedPolymer[:i] + improvedPolymer[i+2:]
				i = 0
			} else {
				i++
			}
		}
		if shortestPolymer > len(improvedPolymer) {
			shortestPolymer = len(improvedPolymer)
		}
	}
	fmt.Println(shortestPolymer)
}
