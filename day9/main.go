package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNextValue(input []int) int {
	diffs := []int{}
	lastValue := 0
	for i := 0; i < len(input)-1; i++ {
		current := input[i]
		next := input[i+1]
		diffs = append(diffs, next-current)
		lastValue = next
	}
	// If all the values are 0, then we are done
	allZeros := true
	for next := range diffs {
		if next != 0 {
			allZeros = false
			break
		}
	}

	nextValue := 0
	if !allZeros {
		nextValue = getNextValue(diffs) + lastValue
	}
	fmt.Printf("Input: %v -> %d\n", input, nextValue)
	return nextValue
}

func getPreviousValue(input []int) int {
	diffs := []int{}
	for i := 0; i < len(input)-1; i++ {
		current := input[i]
		next := input[i+1]
		diffs = append(diffs, next-current)
	}
	// If all the values are 0, then we are done
	allZeros := true
	for next := range diffs {
		if next != 0 {
			allZeros = false
			break
		}
	}

	firstValue := input[0]
	previousValue := 0
	if !allZeros {
		previousValue = firstValue - getPreviousValue(diffs)
	}
	fmt.Printf("Input: %v -> %d\n", input, previousValue)
	return previousValue
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	total := 0

	// read file line by line
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

		nextValue := getNextValue(input)

		total += nextValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return total
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	total := 0

	// read file line by line
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

		nextValue := getPreviousValue(input)

		total += nextValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return total
}

func main() {
	// fmt.Printf("Part 1 answer: %v\n", Part1("data.txt"))
	fmt.Printf("Part 2 answer: %v\n", Part2("data.txt"))
}
