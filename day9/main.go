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

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	total := 0
	// create output1.txt file to write to
	outputFile, _ := os.Create("output1.txt")

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
		outputFile.Write([]byte(strconv.Itoa(nextValue) + "\n"))

		total += nextValue
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
