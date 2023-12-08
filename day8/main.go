package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

type Route struct {
	Start   string
	Current string
	Steps   int
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

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	counter := 0
	nodes := make(map[string]Node)
	instructions := ""
	routes := []Route{}

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
			re := regexp.MustCompile("[^a-zA-Z0-9]+")
			// Remove non characters
			left = re.ReplaceAllString(strings.TrimSpace(parts[0]), "")
			right = re.ReplaceAllString(strings.TrimSpace(parts[1]), "")

			nodes[node] = Node{left, right}

			// If node ends with a A it is a starting node
			if strings.HasSuffix(node, "A") {
				routes = append(routes, Route{node, node, 0})
			}
		}

		counter++
	}

	fmt.Printf("Instructions: %s\n", instructions)
	totalSteps := 0

	// Current node
	currentStep := 0
	maxSteps := len(instructions)

	for {
		totalSteps++
		nextStep := instructions[currentStep]

		for i, route := range routes {
			if nextStep == 'L' {
				routes[i].Current = nodes[route.Current].Left
			} else {
				routes[i].Current = nodes[route.Current].Right
			}

			// Check if we are finished
			if strings.HasSuffix(routes[i].Current, "Z") && routes[i].Steps == 0 {
				fmt.Printf("Route %d finished in %d steps\n", i, totalSteps)
				routes[i].Steps = totalSteps
			}

			if route.Current == "XXX" {
				// error
				log.Fatalf("Error")
			}
		}

		// Check if we are finished
		done := true
		for _, route := range routes {
			if route.Steps == 0 {
				done = false
				break
			}
		}

		if done {
			break
		}

		currentStep = (currentStep + 1) % maxSteps
	}

	values := []int{}
	for _, route := range routes {
		values = append(values, route.Steps)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lcmN(values)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	// a * b = lcm(a, b) * gcd(a, b)
	return (a * b) / gcd(a, b)
}

func lcmN(n []int) int {
	if len(n) == 2 {
		return lcm(n[0], n[1])
	}
	return lcm(n[0], lcmN(n[1:]))
}

func main() {
	fmt.Println("Advent of Code 2019 - Day 8")
	// fmt.Printf("Part 1: %d\n", Part1("data.txt"))
	fmt.Printf("Part 2: %d\n", Part2("data.txt"))
}
