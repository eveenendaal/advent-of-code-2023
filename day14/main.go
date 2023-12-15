package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	total := 0
	columns := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// Move each character into a column
		for i, c := range scanner.Text() {
			if len(columns) <= i {
				columns = append(columns, "")
			}
			columns[i] += string(c)
		}
	}

	fmt.Printf("Columns: %v\n", columns)

	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	return total
}
