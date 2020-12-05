package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, expense)
	}
	sort.Ints(expenses)

	for _, expense := range expenses {
		idx := sort.SearchInts(expenses, 2020-expense)
		if idx < len(expenses) && expenses[idx]+expense == 2020 {
			fmt.Println(expenses[idx] * expense)
			return
		}
	}
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, expense)
	}
	sort.Ints(expenses)

	for i := range expenses {
		smallExpense := expenses[i]
		for j := len(expenses) - 1; j > i; j-- {
			bigExpense := expenses[j]
			if smallExpense+bigExpense > 2020 {
				continue
			}
			idx := sort.SearchInts(expenses, 2020-(smallExpense+bigExpense))
			if idx < len(expenses) && expenses[idx]+smallExpense+bigExpense == 2020 {
				fmt.Println(expenses[idx] * smallExpense * bigExpense)
				return
			}
		}
	}
}
