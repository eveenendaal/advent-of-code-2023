package main

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	direction aoc.Direction
	distance  int
	color     string
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
	verticalBorders := []aoc.Position{start}

	// Loop over instructions
	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			switch instruction.direction {
			case aoc.Up:
				start.Row++
				verticalBorders = append(verticalBorders, start)
			case aoc.Right:
				start.Column++
			case aoc.Down:
				start.Row--
				verticalBorders = append(verticalBorders, start)
			case aoc.Left:
				start.Column--
			}
			borders = append(borders, start)
		}
		fmt.Printf("Line: %v\n", instruction)
	}

	// Find the smallest rectangle that contains all the points
	minX := math.MaxInt16
	maxX := math.MinInt16
	minY := math.MaxInt16
	maxY := math.MinInt16

	for _, border := range borders {
		if border.Column < minX {
			minX = border.Column
		}
		if border.Column > maxX {
			maxX = border.Column
		}
		if border.Row < minY {
			minY = border.Row
		}
		if border.Row > maxY {
			maxY = border.Row
		}
	}

	total := 0

	for y := maxY; y >= minY; y-- {
		outside := true
		for x := minX; x <= maxX; x++ {
			// Check if we've hit a border
			emptySpace := true

			for _, border := range borders {
				if border.Column == x && border.Row == y {
					emptySpace = false
					break
				}
			}

			for _, border := range verticalBorders {
				if border.Column == x && border.Row == y {
					outside = !outside
					break
				}
			}

			if emptySpace && !outside {
				fmt.Print("I")
				total++
			} else if !emptySpace {
				fmt.Print("#")
				total++
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	return total
}
