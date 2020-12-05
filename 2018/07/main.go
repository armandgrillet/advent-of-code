package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	partOne()
}

func partOne() {
	inFile, err := os.Open("input.txt")
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

type Step struct {
	Time         int
	Dependencies []string
}

func partTwo() {
	inFile, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	rawSteps := make(map[string][]string)
	for scanner.Scan() {
		re := regexp.MustCompile("[A-Z]+")
		stepSlice := re.FindAllString(scanner.Text(), -1)
		blocker := stepSlice[1]
		blocked := stepSlice[2]
		if _, ok := rawSteps[blocker]; !ok {
			rawSteps[blocker] = []string{}
		}
		if _, ok := rawSteps[blocked]; !ok {
			rawSteps[blocked] = []string{}
		}
		rawSteps[blocked] = append(rawSteps[blocked], blocker)
	}

	alphabet := make(map[string]int)
	for i := 0; i <= 26; i++ {
		alphabet[string('A'+byte(i))] = 61 + i
	}

	steps := make(map[string]Step)
	for k, v := range rawSteps {
		steps[k] = Step{Time: alphabet[k], Dependencies: v}
	}

	stepsDone := make(map[string]bool)
	stepsInProgress := make(map[string]bool)
	timer := -1
	for len(steps) > 0 {
		for k, v := range steps {
			for i := 0; i < len(v.Dependencies); i++ {
				if stepsDone[v.Dependencies[i]] {
					steps[k] = Step{Time: v.Time, Dependencies: append(v.Dependencies[:i], v.Dependencies[i+1:]...)}
				}
			}
		}

		timer++

		sortedSteps := []string{}
		for k := range steps {
			sortedSteps = append(sortedSteps, k)
		}
		sort.Strings(sortedSteps)

		for k := range steps {
			if !stepsInProgress[k] && len(steps[k].Dependencies) == 0 && len(stepsInProgress) < 4 {
				stepsInProgress[k] = true
			}
		}

		for k := range stepsInProgress {
			steps[k] = Step{Time: steps[k].Time - 1, Dependencies: steps[k].Dependencies}
			if steps[k].Time == 0 {
				stepsDone[k] = true
				delete(steps, k)
				delete(stepsInProgress, k)
			}
		}
	}

	fmt.Println(timer + 1)
}
