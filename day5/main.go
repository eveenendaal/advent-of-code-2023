package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RangeMap struct {
	SourceMin int
	TargetMin int
	Size      int
}

type Step struct {
	RangeMaps []RangeMap
}

// Step List
var Steps []Step = make([]Step, 50)
var Seeds []int = make([]int, 0)

// Handle Step
func HandleStep(step Step, seed int) int {
	output := seed
	// For each RangeMap
	for _, rangeMap := range step.RangeMaps {
		min := rangeMap.SourceMin
		max := rangeMap.SourceMin + rangeMap.Size
		diff := seed - min
		minTarget := rangeMap.TargetMin

		// If seed is between min and max
		if seed >= min && seed <= max {
			// Calculate the output
			output = minTarget + diff
			fmt.Printf("Seed: %d, Min: %d, Max: %d, Diff: %d, MinTarget: %d, Output: %d\n", seed, min, max, diff, minTarget, output)

			break
		}
	}
	return output
}

func HandleSeed(seed int) int {
	output := seed
	// For each step
	for _, step := range Steps {
		if len(step.RangeMaps) > 0 {
			// Handle the step
			output = HandleStep(step, output)
		}
	}
	return output
}

// Create a function to parse the input file
func Part1(filePath string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	stepCounter := 0

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Printf("Step Counter: %d\n", stepCounter)

		// If line contains ":" then it's a step
		if strings.Contains(line, ":") {
			// Parse the step
			if stepCounter == 0 {
				// Split on :
				parts := strings.Split(line, ":")
				// Get part 2, trim spaces, and split on space
				seeds := strings.Split(strings.TrimSpace(parts[1]), " ")
				// Convert to int and store in Seeds
				for _, seed := range seeds {
					seed, _ := strconv.Atoi(seed)
					Seeds = append(Seeds, seed)
				}
			}
			stepCounter += 1
			Steps = append(Steps, Step{})
		} else if len(line) > 0 {
			// split on " "
			parts := strings.Split(line, " ")
			fmt.Printf("Parts: %v, %d\n", parts, len(parts))

			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[0])
			size, _ := strconv.Atoi(parts[2])

			// Create a RangeMap
			rangeMap := RangeMap{
				SourceMin: start,
				TargetMin: end,
				Size:      size,
			}
			Steps[stepCounter].RangeMaps = append(Steps[stepCounter].RangeMaps, rangeMap)
		}
	}

	// Print out seeds
	fmt.Println(Seeds)
	// Print out converters
	//fmt.Println(Steps)
	// Print out the output
	minOutput := math.MaxInt
	for _, seed := range Seeds {
		output := HandleSeed(seed)
		fmt.Printf("Seed: %d, Output: %d\n", seed, output)
		if output < minOutput {
			minOutput = output
		}
	}
	fmt.Println("Min Output:", minOutput)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	Part1("data.txt")
}
