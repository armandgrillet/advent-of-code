package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("02.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ids := []string{}
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	idLen := len(ids[0])
	for i := 0; i < idLen; i++ {
		truncatedIDs := make(map[string]bool, len(ids))
		for _, id := range ids {
			truncatedID := id[:i] + id[i+1:]
			if truncatedIDs[truncatedID] {
				fmt.Println(truncatedID)
				return
			}
			truncatedIDs[truncatedID] = true
		}
	}
}
