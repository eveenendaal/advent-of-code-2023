package day23

import (
	"fmt"
	aoc "github.com/eveenendaal/advent-of-code-2023/aoc"
)

func Part1(filePath string) int {
	characters := aoc.ReadFileToCharacters(filePath)

	var path []aoc.Position
	var slopes []aoc.Position

	for y, row := range characters {
		for x, character := range row {
			switch character {
			case '#':
				// Do nothing
			case '.':
				path = append(path, aoc.Position{Col: x, Row: y})
			case '>':
			case 'v':
				slopes = append(slopes, aoc.Position{Col: x, Row: y})
			default:
				fmt.Printf("Unknown character: %c\n", character)
			}
		}
	}

	return len(characters)
}
