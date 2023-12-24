package day23

import (
	"container/list"
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func (slopeMap SlopeMap) findNeighbors(position aoc.Position) chan aoc.Position {
	Right := aoc.Position{Col: 1, Row: 0}
	Down := aoc.Position{Col: 0, Row: 1}
	Left := aoc.Position{Col: -1, Row: 0}
	Up := aoc.Position{Col: 0, Row: -1}

	directions := []aoc.Position{
		Up,
		Right,
		Left,
		Down,
	}

	neighbors := make(chan aoc.Position)
	go func() {
		defer close(neighbors)

		switch slopeMap.path[position] {
		case '^':
			neighbors <- aoc.Position{position.Col, position.Row - 1}
			return
		case '>':
			neighbors <- aoc.Position{position.Col + 1, position.Row}
			return
		case 'v':
			neighbors <- aoc.Position{position.Col, position.Row + 1}
			return
		case '<':
			neighbors <- aoc.Position{position.Col - 1, position.Row}
			return
		}

		for _, direction := range directions {
			neighbor := aoc.Position{position.Col + direction.Col, position.Row + direction.Row}
			if _, ok := slopeMap.path[neighbor]; !ok {
				continue
			}
			neighbors <- neighbor
		}
	}()
	return neighbors
}

func findTheLongestPath(slopeMap SlopeMap) int {
	start := slopeMap.start
	end := slopeMap.end

	toCheck := list.New()
	toCheck.PushBack([3]interface{}{start, make(map[aoc.Position]struct{}), 0})

	costSoFar := make(map[aoc.Position]int)
	costSoFar[start] = 0

	for toCheck.Len() > 0 {
		element := toCheck.Remove(toCheck.Back()).([3]interface{})
		nextPosition := element[0].(aoc.Position)
		path := element[1].(map[aoc.Position]struct{})

		if nextPosition == end {
			continue
		}

		for newPoint := range slopeMap.findNeighbors(nextPosition) {
			newCost := costSoFar[nextPosition] + 1

			if _, exists := path[newPoint]; exists {
				continue
			}

			if _, exists := costSoFar[newPoint]; !exists || newCost > costSoFar[newPoint] {
				costSoFar[newPoint] = newCost

				newPath := make(map[aoc.Position]struct{})
				for k := range path {
					newPath[k] = struct{}{}
				}
				newPath[newPoint] = struct{}{}

				toCheck.PushFront([3]interface{}{newPoint, newPath, newCost})
			}
		}
	}

	return costSoFar[end]
}

type SlopeMap struct {
	path  map[aoc.Position]rune
	start aoc.Position
	end   aoc.Position
}

func NewSlopeMap(characters [][]rune) SlopeMap {
	var path = make(map[aoc.Position]rune)
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
				path[aoc.Position{Col: x, Row: y}] = character
			case '>':
				position := aoc.Position{Col: x, Row: y}
				path[position] = character
			case 'v':
				position := aoc.Position{Col: x, Row: y}
				path[position] = character
			default:
				fmt.Errorf("Unknown character: %c\n", character)
			}
		}
	}

	return SlopeMap{path: path, start: start, end: end}
}

type Progress struct {
	position aoc.Position
	path     map[aoc.Position]struct{}
}

func Part1(filePath string) int {
	characters := aoc.ReadFileToCharacters(filePath)
	slopeMap := NewSlopeMap(characters)
	fmt.Printf("Start: %v, Stop %v\n", slopeMap.start, slopeMap.end)
	return findTheLongestPath(slopeMap)
}
