package day21

import (
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

func nextStep(steps int, start aoc.Position, gardens []aoc.Position, visited []aoc.Position) []aoc.Position {
	if steps == 0 {
		return visited
	}

	neighbors := findNeighbors(start, gardens)
	for _, neighbor := range neighbors {
		if !aoc.ContainsPosition(visited, neighbor) {
			visited = append(visited, neighbor)
		}
	}

	for _, neighbor := range neighbors {
		visited = nextStep(steps-1, neighbor, gardens, visited)
	}

	return visited
}

func Part1(filePath string, steps int) int {
	start, gardens := parseInput(filePath)
	visited := nextStep(steps-1, start, gardens, []aoc.Position{})
	return len(visited)
}
