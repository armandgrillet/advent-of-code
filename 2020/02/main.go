package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	totalMatches := 0
	rowRegExp := regexp.MustCompile(`(.*?)-(.*?) (.*?): (.*?)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := rowRegExp.FindStringSubmatch(scanner.Text())
		min, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		letter := match[3]
		password := match[4]

		letterRexExp := regexp.MustCompile(letter)
		matches := len(letterRexExp.FindAllStringIndex(password, -1))
		if matches >= min && matches <= max {
			totalMatches++
		}
	}
	fmt.Println(totalMatches)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	totalMatches := 0
	rowRegExp := regexp.MustCompile(`(.*?)-(.*?) (.*?): (.*?)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := rowRegExp.FindStringSubmatch(scanner.Text())
		firstPos, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		secondPos, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		letter := match[3]
		password := match[4]
		if (string(password[firstPos-1]) == letter && string(password[secondPos-1]) != letter) ||
			(string(password[firstPos-1]) != letter && string(password[secondPos-1]) == letter) {
			totalMatches++
		}
	}
	fmt.Println(totalMatches)
}
