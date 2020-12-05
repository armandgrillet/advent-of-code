package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	validPassports := 0
	mandatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // cid is ignored

	scanner := bufio.NewScanner(file)
	passport := ""
	line := ""
	valid := true

	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			passport += line + " "
		} else {
			for _, field := range mandatoryFields {
				if strings.Contains(passport, field+":") == false {
					valid = false
					break
				}
			}
			if valid {
				validPassports++
			}
			passport = ""
			valid = true
		}
	}
	fmt.Println(validPassports)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	validPassports := 0
	mandatoryExpressions := []*regexp.Regexp{
		regexp.MustCompile(`byr:(19[2-9][0-9]|200[0-2]) `),
		regexp.MustCompile(`iyr:20(1[0-9]|20) `),
		regexp.MustCompile(`eyr:20(2[0-9]|30) `),
		regexp.MustCompile(`hgt:(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in) `),
		regexp.MustCompile(`hcl:#[0-9a-f]{6} `),
		regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth) `),
		regexp.MustCompile(`pid:[0-9]{9} `),
	}

	scanner := bufio.NewScanner(file)
	passport := ""
	line := ""
	valid := true

	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			passport += line + " "
		} else {
			for _, expression := range mandatoryExpressions {
				if expression.MatchString(passport) == false {
					valid = false
					break
				}
			}
			if valid {
				validPassports++
			}
			passport = ""
			valid = true
		}
	}
	fmt.Println(validPassports)
}
