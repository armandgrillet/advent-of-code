package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("04.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	events := []string{}
	for scanner.Scan() {
		events = append(events, scanner.Text())
	}
	sort.Strings(events)

	guards := make(map[int][]int)
	var feltAsleep bool
	var lastGuard, feltAsleepMin int
	for _, event := range events {
		// [1518-06-07 23:59] Guard #61 => ["1518", "06", "07", "23", "59", "61"]
		re := regexp.MustCompile("[0-9]+")
		rawGuardSlice := re.FindAllString(event, -1)

		guardSlice := []int{}
		for _, strVal := range rawGuardSlice {
			val, err := strconv.Atoi(strVal)
			if err != nil {
				panic(err)
			}
			guardSlice = append(guardSlice, val)
		}

		// Update the outstanding times.
		if guardSlice[3] != 0 {
			guardSlice[4] = 0
		}
		// We have an ID in the line
		if len(guardSlice) == 6 {
			lastGuard = guardSlice[5]
			if _, ok := guards[lastGuard]; !ok {
				guards[lastGuard] = make([]int, 60)
			}
		} else {
			if strings.Contains(event, "asleep") {
				feltAsleep = true
				feltAsleepMin = guardSlice[4]
			} else if strings.Contains(event, "up") && feltAsleep {
				for i := feltAsleepMin; i < guardSlice[4]; i++ {
					guards[lastGuard][i]++
				}
				feltAsleep = false
			}
		}
	}

	var sleepyGuard, maxMinsAsleep int
	for id, schedule := range guards {
		minsAsleep := 0
		for _, minute := range schedule {
			minsAsleep += minute
		}

		if minsAsleep > maxMinsAsleep {
			sleepyGuard = id
			maxMinsAsleep = minsAsleep
		}
	}

	var bestMinute, maxTimesAsleep int
	for i, minute := range guards[sleepyGuard] {
		if minute > maxTimesAsleep {
			bestMinute = i
			maxTimesAsleep = minute
		}
	}
	fmt.Println(sleepyGuard * bestMinute)
}
