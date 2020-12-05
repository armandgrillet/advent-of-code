package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	id   int
	left int
	top  int
	wide int
	tall int
}

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	claims := []claim{}
	var maxLength, maxHeight int
	for scanner.Scan() {
		// #123 @ 3,2: 5x4 => ["123", "3", "2", "5", "4"]
		re := regexp.MustCompile("[0-9]+")
		rawClaimSlice := re.FindAllString(scanner.Text(), -1)

		claimSlice := []int{}
		for _, strVal := range rawClaimSlice {
			val, err := strconv.Atoi(strVal)
			if err != nil {
				panic(err)
			}
			claimSlice = append(claimSlice, val)
		}

		c := claim{
			id:   claimSlice[0],
			left: claimSlice[1],
			top:  claimSlice[2],
			wide: claimSlice[3],
			tall: claimSlice[4],
		}

		if c.left+c.wide > maxLength {
			maxLength = c.left + c.wide
		}
		if c.top+c.tall > maxHeight {
			maxHeight = c.top + c.tall
		}

		claims = append(claims, c)
	}

	fabric := make([][]int, maxLength)
	for i := range fabric {
		fabric[i] = make([]int, maxHeight)
	}

	for _, c := range claims {
		for i := c.left; i < c.left+c.wide; i++ {
			for j := c.top; j < c.top+c.tall; j++ {
				fabric[i][j]++
			}
		}
	}

	overclaimedInches := 0
	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				overclaimedInches++
			}
		}
	}

	fmt.Println(overclaimedInches)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	claims := []claim{}
	var maxLength, maxHeight int
	for scanner.Scan() {
		// #123 @ 3,2: 5x4 => ["123", "3", "2", "5", "4"]
		re := regexp.MustCompile("[0-9]+")
		rawClaimSlice := re.FindAllString(scanner.Text(), -1)

		claimSlice := []int{}
		for _, strVal := range rawClaimSlice {
			val, err := strconv.Atoi(strVal)
			if err != nil {
				panic(err)
			}
			claimSlice = append(claimSlice, val)
		}

		c := claim{
			id:   claimSlice[0],
			left: claimSlice[1],
			top:  claimSlice[2],
			wide: claimSlice[3],
			tall: claimSlice[4],
		}

		if c.left+c.wide > maxLength {
			maxLength = c.left + c.wide
		}
		if c.top+c.tall > maxHeight {
			maxHeight = c.top + c.tall
		}

		claims = append(claims, c)
	}

	fabric := make([][]int, maxLength)
	for i := range fabric {
		fabric[i] = make([]int, maxHeight)
	}

	for _, c := range claims {
		for i := c.left; i < c.left+c.wide; i++ {
			for j := c.top; j < c.top+c.tall; j++ {
				fabric[i][j]++
			}
		}
	}

	for _, c := range claims {
		overlapping := false
		for i := c.left; i < c.left+c.wide; i++ {
			for j := c.top; j < c.top+c.tall; j++ {
				if fabric[i][j] > 1 {
					overlapping = true
				}
			}
		}
		if !overlapping {
			fmt.Println(c.id)
			return
		}
	}
}
