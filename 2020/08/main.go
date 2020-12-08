package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	cmd    string
	change int
}

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	instructions := []instruction{}

	r := regexp.MustCompile(`(\+|-)\d+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.Atoi(r.FindAllString(scanner.Text(), 1)[0])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{string(scanner.Text())[:3], change})
	}

	res, _ := runInstructions(instructions)
	fmt.Println(res)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	instructions := []instruction{}

	r := regexp.MustCompile(`(\+|-)\d+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.Atoi(r.FindAllString(scanner.Text(), 1)[0])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{string(scanner.Text())[:3], change})
	}

	for i := 0; i < len(instructions); i++ {
		if instructions[i].cmd == "jmp" {
			instructions[i].cmd = "nop"
		} else if instructions[i].cmd == "nop" {
			instructions[i].cmd = "jmp"
		}

		res, err := runInstructions(instructions)
		if err == nil {
			fmt.Println(res)
			return
		}

		if instructions[i].cmd == "jmp" {
			instructions[i].cmd = "nop"
		} else if instructions[i].cmd == "nop" {
			instructions[i].cmd = "jmp"
		}
	}
}

func runInstructions(instructions []instruction) (int, error) {
	pos := 0
	acc := 0
	posAndAcc := make(map[int]bool)
	posChange := 1
	accChange := 0
	for pos >= 0 && pos < len(instructions) && posAndAcc[pos] == false {
		posAndAcc[pos] = true
		posChange = 1
		accChange = 0

		instruction := instructions[pos]
		switch instruction.cmd {
		case "acc":
			accChange = instruction.change
		case "jmp":
			posChange = instruction.change
		}

		if posAndAcc[pos+posChange] {
			return acc, errors.New("infinite loop detected")
		}
		pos += posChange
		acc += accChange
	}

	return acc, nil
}
