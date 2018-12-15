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

	lenPolymers := make(chan int, 26)
	for letter := 0; letter < 26; letter++ {
		go func(letter int, polymer string, lenPolymers chan<- int) {
			producedPolymer := strings.Replace(polymer, string('A'+letter), "", -1)
			producedPolymer = strings.Replace(producedPolymer, string('a'+letter), "", -1)
			i := 0
			for i < len(producedPolymer)-1 {
				if producedPolymer[i]+32 == producedPolymer[i+1] || producedPolymer[i]-32 == producedPolymer[i+1] {
					producedPolymer = producedPolymer[:i] + producedPolymer[i+2:]
					i = 0
				} else {
					i++
				}
			}
			lenPolymers <- len(producedPolymer)
		}(letter, polymer, lenPolymers)
	}

	shortestPolymer := len(polymer)
	for i := 0; i < 26; i++ {
		lenPolymer := <-lenPolymers
		if shortestPolymer > lenPolymer {
			shortestPolymer = lenPolymer
		}
	}
	fmt.Println(shortestPolymer)
}
