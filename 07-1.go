package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inFile, err := os.Open("07.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	steps := make(map[string][]string)
	for scanner.Scan() {
		re := regexp.MustCompile("[A-Z]+")
		stepSlice := re.FindAllString(scanner.Text(), -1)
		step := stepSlice[1]
		dep := stepSlice[2]
		if _, ok := steps[step]; !ok {
			steps[step] = []string{}
		}
		if _, ok := steps[dep]; !ok {
			steps[dep] = []string{}
		}
		steps[step] = append(steps[step], dep)
	}

	alphabet := make([]string, 26)
	for i := range alphabet {
		alphabet[i] = string('A' + byte(i))
	}

	for len(steps) > 0 {
		for i := 0; i < len(alphabet); i++ {
			letter := alphabet[i]
			existingLetter, firstLetter := false, true
			for step, deps := range steps {
				if step == letter {
					existingLetter = true
				}
				for _, dep := range deps {
					if dep == letter {
						firstLetter = false
					}
				}
			}

			if existingLetter && firstLetter {
				fmt.Print(letter)
				delete(steps, letter)
				i = -1
			}
		}
	}
}
