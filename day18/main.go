package main

import (
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"strconv"
	"strings"
)

type Instruction struct {
	direction aoc.Direction
	distance  int
	color     string
}

func findAreaShoelace(nodes []aoc.Position) int {
	area := 0

	for i := 0; i < len(nodes); i++ {
		pointA, pointB := nodes[i], nodes[(i+1)%(len(nodes))]
		area += (pointA.Column * pointB.Row) - (pointB.Column * pointA.Row) + max(aoc.IntAbs(pointA.Column-pointB.Column), aoc.IntAbs(pointA.Row-pointB.Row))
	}

	return area / 2
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

func Solve(filepath string) int {
	instructions := make([]Instruction, 0)

	lines := aoc.ReadFileToLines(filepath)
	for _, line := range lines {
		// splits on " "
		parts := strings.Split(line, " ")

		// Handle Color
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
