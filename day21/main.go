package day21

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func parseInput(filePath string) (aoc.Position, []aoc.Position) {
	characters := aoc.ReadFileToCharacters(filePath)
	var start aoc.Position
	var gardens []aoc.Position

	for y, line := range characters {
		for x, character := range line {
			switch character {
			case 'S':
				start = aoc.Position{Col: x, Row: y}
				gardens = append(gardens, aoc.Position{Col: x, Row: y})
			case '.':
				gardens = append(gardens, aoc.Position{Col: x, Row: y})
			}
		}
	}

	return start, gardens
}

func isNeighbor(p aoc.Position, other aoc.Position) bool {
	if p.Col == other.Col && p.Row == other.Row {
		return false
	}
	if p.Col == other.Col && (p.Row == other.Row-1 || p.Row == other.Row+1) {
		return true
	}
	if p.Row == other.Row && (p.Col == other.Col-1 || p.Col == other.Col+1) {
		return true
	}
	return false
}

func findNeighbors(position aoc.Position, positions []aoc.Position) []aoc.Position {
	neighbors := make([]aoc.Position, 0)
	for _, candidate := range positions {
		if isNeighbor(position, candidate) {
			neighbors = append(neighbors, candidate)
		}
	}
	return neighbors
}

func nextStep(steps int, edges []aoc.Position, gardens []aoc.Position) []aoc.Position {
	newEdges := make([]aoc.Position, 0)

	for _, edge := range edges {
		neighbors := findNeighbors(edge, gardens)
		for _, neighbor := range neighbors {
			if !aoc.ContainsPosition(newEdges, neighbor) {
				newEdges = append(newEdges, neighbor)
			}
		}
	}

	// printGardens(steps, gardens, newEdges)
	if steps == 0 {
		return newEdges
	} else {
		return nextStep(steps-1, newEdges, gardens)
	}
}

func printGardens(steps int, gardens []aoc.Position, visited []aoc.Position) {
	maxY := 0
	maxX := 0

	fmt.Printf("After %d steps:\n", 6-steps)

	for _, garden := range gardens {
		if garden.Row > maxY {
			maxY = garden.Row
		}
		if garden.Col > maxX {
			maxX = garden.Col
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if aoc.ContainsPosition(visited, aoc.Position{Col: x, Row: y}) {
				print("O")
			} else if aoc.ContainsPosition(gardens, aoc.Position{Col: x, Row: y}) {
				print(".")
			} else {
				print("#")
			}
		}
		println()
	}
	println()
}

func Part1(filePath string, steps int) int {
	start, gardens := parseInput(filePath)
	edges := nextStep(steps-1, []aoc.Position{start}, gardens)
	return len(edges)
}
