package day23

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func (slopeMap SlopeMap) findNeighbors(position aoc.Position) []aoc.Position {
	directions := []aoc.Position{
		{Col: 0, Row: -1},
		{Col: 1, Row: 0},
		{Col: 0, Row: 1},
		{Col: -1, Row: 0},
	}

	var neighbors []aoc.Position
	for _, direction := range directions {
		neighbor := aoc.Position{Col: position.Col + direction.Col, Row: position.Row + direction.Row}
		if slopeMap.isPath(neighbor) && !slopeMap.isUphill(neighbor, direction) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (slopeMap SlopeMap) isUphill(position aoc.Position, direction aoc.Position) bool {
	if slopeMap.slopes[position] == aoc.Right && direction.Col == -1 {
		return true
	} else if slopeMap.slopes[position] == aoc.Down && direction.Row == -1 {
		return true
	}
	return false
}

func (slopeMap SlopeMap) isPath(position aoc.Position) bool {
	for _, path := range slopeMap.path {
		if path == position {
			return true
		}
	}
	return false
}

func (slopeMap SlopeMap) Step(visited []aoc.Position, position aoc.Position) int {
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
		return len(visited)
	} else {
		for _, neighbor := range unvisited {
			slopeMap.Step(visited, neighbor)
		}
	}

}

type SlopeMap struct {
	path   []aoc.Position
	slopes map[aoc.Position]aoc.Direction
	start  aoc.Position
	end    aoc.Position
}

func NewSlopeMap(characters [][]rune) SlopeMap {
	var path []aoc.Position
	var slopes = make(map[aoc.Position]aoc.Direction)

	for y, row := range characters {
		for x, character := range row {
			switch character {
			case '#':
				// Do nothing
			case '.':
				path = append(path, aoc.Position{Col: x, Row: y})
			case '>':
				position := aoc.Position{Col: x, Row: y}
				path = append(path, position)
				slopes[position] = aoc.Right
			case 'v':
				position := aoc.Position{Col: x, Row: y}
				path = append(path, position)
				slopes[position] = aoc.Down
			default:
				fmt.Errorf("Unknown character: %c\n", character)
			}
		}
	}

	return SlopeMap{path: path, slopes: slopes}
}

func Part1(filePath string) int {
	characters := aoc.ReadFileToCharacters(filePath)
	slopeMap := NewSlopeMap(characters)
	slopeMap.Step([]aoc.Position{}, slopeMap.path[0])
	return len(characters)
}
