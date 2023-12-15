package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func reversedMatch(a, b []string) bool {
	for i := range a {
		if a[i] != b[len(b)-1-i] {
			return false
		}
	}
	return true
}

func transpose(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	result := make([]string, len(strings[0]))
	for _, s := range strings {
		for i, r := range s {
			result[i] = result[i] + string(r)
		}
	}

	return result
}

func checkBlock(pattern []string) bool {
	if len(pattern)%2 != 0 {
		panic("Pattern must be even in length")
	}
	top := 0
	bottom := len(pattern)
	middle := bottom / 2

	topHalf := sort.StringSlice(pattern[top:middle])
	bottomHalf := sort.StringSlice(pattern[middle:bottom])

	return reversedMatch(topHalf, bottomHalf)
}

func handlePattern(horizontalPattern []string) int {
	// Find vertical match
	totalVerticalLines := len(horizontalPattern)
	start := totalVerticalLines - (totalVerticalLines % 2) // It must be an even number
	total := 0

	horizontalLineFound := false
	verticalLineFound := false

	for i := start; i > 1; i -= 2 {
		// Check from Top
		block := horizontalPattern[0:i]
		if checkBlock(block) {
			lines := i / 2
			fmt.Printf("Vertical Match: %d -> %v\n", lines, block)
			total += lines * 100
			horizontalLineFound = true
			break
		}

		// Check from Bottom
		block = horizontalPattern[totalVerticalLines-i : totalVerticalLines]
		if checkBlock(block) {
			lines := totalVerticalLines - (i / 2)
			fmt.Printf("Vertical Match: %d -> %v\n", lines, block)
			total += lines * 100
			horizontalLineFound = true
			break
		}
	}

	// Build horizontal lines
	verticalPattern := transpose(horizontalPattern)
	totalHorizontalLines := len(verticalPattern)
	start = totalHorizontalLines - (totalHorizontalLines % 2) // It must be an even number

	for i := start; i > 1; i -= 2 {
		// Check from Top
		block := verticalPattern[0:i]
		if checkBlock(block) {
			lines := i / 2
			fmt.Printf("Horizontal Match: %d -> %v\n", lines, block)
			total += lines
			verticalLineFound = true
			break
		}

		// Check from Bottom
		block = verticalPattern[totalHorizontalLines-i : totalHorizontalLines]
		if checkBlock(block) {
			lines := totalHorizontalLines - (i / 2)
			fmt.Printf("Horizontal Match: %d -> %v\n", lines, block)
			total += lines
			verticalLineFound = true
			break
		}
	}

	if !horizontalLineFound && !verticalLineFound {
		panic("No match found")
	}

	return total
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

func main() {
	fmt.Println("Advent of Code 2017 - Day 13")
	fmt.Println("Part 1:", Part1("data.txt"))
}
