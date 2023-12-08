package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	counter := 0
	nodes := make(map[string]Node)
	instructions := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if counter == 0 {
			instructions = line
		} else if len(line) > 0 {
			var node, left, right string

			// split on "="
			parts := strings.Split(line, "=")
			node = strings.TrimSpace(parts[0])
			parts = strings.Split(parts[1], ",")
			// Create regex for non characters
			re := regexp.MustCompile("[^a-zA-Z]+")
			// Remove non characters
			left = re.ReplaceAllString(strings.TrimSpace(parts[0]), "")
			right = re.ReplaceAllString(strings.TrimSpace(parts[1]), "")

			nodes[node] = Node{left, right}
			fmt.Printf("Name: %s, Left: %s, Right: %s\n", node, left, right)
		}

		counter++
	}

	fmt.Printf("Instructions: %s\n", instructions)
	totalSteps := 0

	// Current node
	current := "AAA"
	currentStep := 0
	maxSteps := len(instructions)

	for {
		totalSteps++
		nextStep := instructions[currentStep]

		if nextStep == 'L' {
			current = nodes[current].Left
		} else {
			current = nodes[current].Right
		}

		// Check if we are finished
		if current == "ZZZ" {
			break
		}
		currentStep = (currentStep + 1) % maxSteps
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return totalSteps
}
