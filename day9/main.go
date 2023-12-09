package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNextValue(input []int) int {
	fmt.Printf("Input: %v\n", input)

	diffs := []int{}
	lastValue := 0
	for i := 0; i < len(input)-1; i++ {
		current := input[i]
		next := input[i+1]
		diffs = append(diffs, next-current)
		lastValue = next
	}
	// If all the values are 0, then we are done

	if lastValue == 0 {
		return 0
	} else {
		return getNextValue(diffs) + lastValue
	}
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split line on spaces
		rawInput := strings.Split(line, " ")
		input := []int{}
		// convert from string to int
		for _, v := range rawInput {
			intValue, _ := strconv.Atoi(v)
			input = append(input, intValue)
		}

		total += getNextValue(input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return total
}

func main() {
	result := Part1("data.txt")
	fmt.Printf("Part 1 answer: %v\n", result)
}
