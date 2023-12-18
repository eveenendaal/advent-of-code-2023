package main

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	direction aoc.Direction
	distance  int
	color     string
}

func floodFill(border []aoc.Position, start aoc.Position) int {
	trench := make(map[aoc.Position]bool)

	// Add border to trench
	for _, position := range border {
		trench[position] = true
	}

	queue := []aoc.Position{start}

	for {
		if len(queue) == 0 {
			break
		}

		// Get next
		next := queue[0]

		// get edges
		edges := []aoc.Position{{next.Column, next.Row - 1}, {next.Column, next.Row + 1}, {next.Column - 1, next.Row}, {next.Column + 1, next.Row}}

		for _, edge := range edges {
			if _, contains := trench[edge]; !contains {
				trench[edge] = true
				queue = append(queue, edge)
			}
		}

		// remove next from queue
		queue = queue[1:]
	}

	// Convert trench to slice
	return len(trench)
}

func Part1(filepath string) int {
	instructions := make([]Instruction, 0)

	lines := aoc.ReadFileToLines(filepath)
	for _, line := range lines {
		// splits on " "
		parts := strings.Split(line, " ")

		// Handle Direction
		direction := aoc.Up
		switch parts[0] {
		case "R":
			direction = aoc.Right
		case "L":
			direction = aoc.Left
		case "U":
			direction = aoc.Up
		case "D":
			direction = aoc.Down
		}

		// Handle Distance
		distance, _ := strconv.Atoi(parts[1])

		// Create Instruction
		instruction := Instruction{
			direction: direction,
			distance:  distance,
			color:     parts[2],
		}

		// Add to instructions
		instructions = append(instructions, instruction)
	}

	// Create Start
	start := aoc.Position{Column: 0, Row: 0}
	borders := []aoc.Position{start}

	// Loop over instructions
	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			switch instruction.direction {
			case aoc.Up:
				start.Row++
			case aoc.Right:
				start.Column++
			case aoc.Down:
				start.Row--
			case aoc.Left:
				start.Column--
			}
			borders = append(borders, start)
		}
	}

	area := floodFill(borders, aoc.Position{1, -1})
	return area
}

func findAreaShoelace(nodes []aoc.Position) int {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0

	for i := 0; i < len(nodes); i++ {
		current := nodes[i]
		next := nodes[(i+1)%(len(nodes))]

		area += (current.Column * next.Row) - (next.Column * current.Row) + max(
			aoc.IntAbs(current.Column-next.Column),
			aoc.IntAbs(current.Row-next.Row),
		)
	}

	return area / 2
}

func Part2(filepath string) int {
	instructions := make([]Instruction, 0)

	lines := aoc.ReadFileToLines(filepath)
	for _, line := range lines {
		// splits on " "
		parts := strings.Split(line, " ")

		// Ignore 1 and 2
		lineInput := parts[2]

		// remove non-hex characters
		lineInput = regexp.MustCompile("[^0-9a-fA-F]+").ReplaceAllString(lineInput, "")

		// Get last character
		directionString := string(lineInput[len(lineInput)-1])
		// Get the rest
		distanceString := lineInput[:len(lineInput)-1]
		// convert from hex to int
		distance, _ := strconv.ParseInt(distanceString, 16, 64)

		// Handle Direction
		direction := aoc.Up
		switch directionString {
		case "0":
			direction = aoc.Right
		case "1":
			direction = aoc.Up
		case "2":
			direction = aoc.Left
		case "3":
			direction = aoc.Down
		}

		instructions = append(instructions, Instruction{
			direction: direction,
			distance:  int(distance),
		})
	}

	for _, instruction := range instructions {
		fmt.Printf("Direction: %d, Distance: %d\n", instruction.direction, instruction.distance)
	}

	// Create Start
	start := aoc.Position{Column: 0, Row: 0}
	borders := []aoc.Position{start}

	// Loop over instructions
	for _, instruction := range instructions {
		switch instruction.direction {
		case aoc.Up:
			start.Row += instruction.distance
		case aoc.Right:
			start.Column += instruction.distance
		case aoc.Down:
			start.Row -= instruction.distance
		case aoc.Left:
			start.Column -= instruction.distance
		}
		borders = append(borders, start)
	}

	// Convert borders to points
	area := findAreaShoelace(borders) + 1
	return area
}
