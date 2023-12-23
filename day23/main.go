package day23

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

type Rule int

const (
	AllowAll Rule = iota
	AllowRight
	AllowDown
)

func (slopeMap SlopeMap) findNeighbors(position aoc.Position) []aoc.Position {
	Right := aoc.Position{Col: 1, Row: 0}
	Down := aoc.Position{Col: 0, Row: 1}
	Left := aoc.Position{Col: -1, Row: 0}
	Up := aoc.Position{Col: 0, Row: -1}

	directions := []aoc.Position{
		Right,
		Left,
		Down,
		Up,
	}

	var neighbors []aoc.Position
	for _, direction := range directions {
		neighbor := aoc.Position{Col: position.Col + direction.Col, Row: position.Row + direction.Row}
		if _, ok := slopeMap.path[neighbor]; ok {
			switch slopeMap.path[neighbor] {
			case AllowAll:
				neighbors = append(neighbors, neighbor)
			case AllowRight:
				if direction == Right {
					neighbors = append(neighbors, neighbor)
				}
			case AllowDown:
				if direction == Down {
					neighbors = append(neighbors, neighbor)
				}
			}
		}

	}
	return neighbors
}

func (slopeMap SlopeMap) Step(visited []aoc.Position, position aoc.Position) []int {
	var result []int
	// Mark position as visited
	visited = append(visited, position)

	// Find neighbors
	neighbors := slopeMap.findNeighbors(position)

	// Filter visited neighbors
	var unvisited []aoc.Position
	for _, neighbor := range neighbors {
		if !aoc.ContainsPosition(visited, neighbor) {
			unvisited = append(unvisited, neighbor)
		}
	}

	// Recursively step through unvisited neighbors
	if len(unvisited) == 0 {
		if position == slopeMap.end {
			return []int{len(visited)}
		} else {
			return []int{}
		}
	} else {
		for _, neighbor := range unvisited {
			paths := slopeMap.Step(visited, neighbor)
			result = append(result, paths...)
		}
		return result
	}
}

type SlopeMap struct {
	path  map[aoc.Position]Rule
	start aoc.Position
	end   aoc.Position
}

func NewSlopeMap(characters [][]rune) SlopeMap {
	var path = make(map[aoc.Position]Rule)
	var start aoc.Position
	var end aoc.Position

	for y, row := range characters {
		for x, character := range row {
			switch character {
			case '#':
				// Do nothing
			case '.':
				if y == 0 {
					start = aoc.Position{Col: x, Row: y}
				} else if y == len(characters)-1 {
					end = aoc.Position{Col: x, Row: y}
				}
				path[aoc.Position{Col: x, Row: y}] = AllowAll
			case '>':
				position := aoc.Position{Col: x, Row: y}
				path[position] = AllowRight
			case 'v':
				position := aoc.Position{Col: x, Row: y}
				path[position] = AllowDown
			default:
				fmt.Errorf("Unknown character: %c\n", character)
			}
		}
	}

	return SlopeMap{path: path, start: start, end: end}
}

func Part1(filePath string) int {
	characters := aoc.ReadFileToCharacters(filePath)
	slopeMap := NewSlopeMap(characters)
	results := slopeMap.Step([]aoc.Position{}, slopeMap.start)
	maxSteps := 0
	for _, result := range results {
		if result > maxSteps {
			maxSteps = result
		}
	}
	return maxSteps
}
