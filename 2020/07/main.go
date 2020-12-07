package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`\w+ \w+ bag`)
	bagsAndContainers := make(map[string][]string)
	var matches []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches = r.FindAllString(scanner.Text(), -1)
		if len(matches) > 2 || (len(matches) == 2 && matches[1] != "no other bag") {
			for i := 1; i < len(matches); i++ {
				bagsAndContainers[matches[i]] = append(bagsAndContainers[matches[i]], matches[0])
			}
		}
	}

	fmt.Println(searchContainers("shiny gold bag", bagsAndContainers, make(map[string]bool)))
}

func searchContainers(bag string, bagsAndContainers map[string][]string, containers map[string]bool) int {
	for _, container := range bagsAndContainers[bag] {
		if containers[container] == false {
			containers[container] = true
			searchContainers(container, bagsAndContainers, containers)
		}
	}
	return len(containers)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`(\w+ \w+|\d+ \w+ \w+) bag`)
	bagsAndContent := map[string]map[string]int{}
	var matches []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches = r.FindAllString(scanner.Text(), -1)
		if len(matches) > 2 || (len(matches) == 2 && matches[1] != "no other bag") {
			bagsAndContent[matches[0]] = map[string]int{}
			for i := 1; i < len(matches); i++ {
				countAndBags := strings.SplitN(matches[i], " ", 2)
				count, err := strconv.Atoi(countAndBags[0])
				if err != nil {
					panic(err)
				}
				bagsAndContent[matches[0]][countAndBags[1]] = count // "dark red bag": { "dark orange bag" : 2, "pale blue bag" : 5}
			}
		}
	}

	// The -1 is because we want to know how many are in the shiny gold bag thus
	// we eclude the siny gold bag.
	fmt.Println(countBagsInBag("shiny gold bag", bagsAndContent) - 1)
}

func countBagsInBag(bag string, bagsAndContent map[string]map[string]int) int {
	bagsContained := 1
	for bag, count := range bagsAndContent[bag] {
		bagsContained += count * countBagsInBag(bag, bagsAndContent)
	}
	return bagsContained
}
