package main

import (
	"bufio"
	"fmt"
	"os"
)

func slicesMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkBlock(pattern []string) bool {
	if len(pattern)%2 != 0 {
		panic("Pattern must be even in length")
	}
	top := 0
	bottom := len(pattern)
	middle := bottom / 2

	topHalf := pattern[top:middle]
	bottomHalf := pattern[middle:bottom]

	// Reverse the bottom half
	for i, j := 0, len(bottomHalf)-1; i < j; i, j = i+1, j-1 {
		bottomHalf[i], bottomHalf[j] = bottomHalf[j], bottomHalf[i]
	}

	return slicesMatch(topHalf, bottomHalf)
}

func handlePattern(pattern []string) int {
	// Find vertical match
	totalVerticalLines := len(pattern)
	start := totalVerticalLines - (totalVerticalLines % 2) // It must be an even number

	for i := start; i > 1; i -= 2 {
		// Top First
		// Get i number of rows from the top
		n := i / 2

		fmt.Printf("i: %v, n: %v\n", i, n)

		block := pattern[0:i]
		if checkBlock(block) {
			fmt.Printf("Match: %d -> %v\n", i, block)
		}

		block = pattern[i:totalVerticalLines]
		if checkBlock(block) {
			fmt.Printf("Match: %d -> %v\n", i, block)
		}

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
