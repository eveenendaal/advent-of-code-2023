package main

import (
	"bufio"
	"fmt"
	"os"
)

func handlePattern(pattern []string) int {
	// Find vertical match
	totalVerticalLines := len(pattern)
	start := totalVerticalLines - (totalVerticalLines % 2) // It must be an even number

	for i := start; i > 1; i -= 2 {
		// Top First
		// Get i number of rows from the top
		n := i / 2
		top := pattern[0:n]
		bottom := pattern[n:i]
		fmt.Printf("Top: %v\n", top)
		fmt.Printf("Bottom: %v\n", bottom)

		// Get i number of rows from the bottom

		// Bottom Next

	}

	fmt.Println(pattern)
	// Build hortizontal lines
	// Find horizontal match
	return 0
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			total += handlePattern(lines)
			lines = []string{}
		}
	}
	total += handlePattern(lines)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	return total
}
