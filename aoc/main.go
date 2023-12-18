package base

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

// ReadFileToLines reads a file into a slice of strings
func ReadFileToLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// ReadFileToCharacters reads a file into a 2D array of runes
func ReadFileToCharacters(filePath string) [][]rune {
	lines := ReadFileToLines(filePath)
	characters := make([][]rune, len(lines))
	for y, line := range lines {
		for x, char := range line {
			if characters[y] == nil {
				characters[y] = make([]rune, len(line))
			}
			characters[y][x] = char
		}
	}
	return characters
}

// ReadFileToInt reads a file into a slice of ints
func ReadFileToInt(filePath string) [][]int {
	lines := ReadFileToLines(filePath)
	results := make([][]int, len(lines))
	for y, line := range lines {
		for x, char := range line {
			if results[y] == nil {
				results[y] = make([]int, len(line))
			}
			results[y][x], _ = strconv.Atoi(string(char))
		}
	}
	return results
}

type Position struct {
	Col int
	Row int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// IntAbs Absolute value of an integer
func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// FindInclusiveAreaShoelace Finds the area of a polygon using the shoelace formula
func FindInclusiveAreaShoelace(nodes []Position) int {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0

	for i := 0; i < len(nodes); i++ {
		current := nodes[i]
		next := nodes[(i+1)%(len(nodes))]

		area += (current.Col * next.Row) - (next.Col * current.Row) + max(
			IntAbs(current.Col-next.Col),
			IntAbs(current.Row-next.Row),
		)
	}

	return (area / 2) + 1
}

// FindInteriorAreaShoelace Finds the area inside a polygon using the shoelace formula
func FindInteriorAreaShoelace(original []Position) int {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0

	nodes := append([]Position(nil), original...)

	// Adjust the vertices inward by 1 pixel
	for i := range nodes {
		nodes[i].Col = max(0, nodes[i].Col-1)
		nodes[i].Row = max(0, nodes[i].Row-1)
	}

	for i := 0; i < len(nodes); i++ {
		current := nodes[i]
		next := nodes[(i+1)%(len(nodes))]

		area += (current.Col * next.Row) - (next.Col * current.Row)
	}

	return area / 2
}

// ContainsPosition Checks if a slice of positions contains a position
func ContainsPosition(points []Position, point Position) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

// PrintShape Prints a shape to the console
func PrintShape(points []Position) {
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for _, point := range points {
		minX = min(minX, point.Col)
		minY = min(minY, point.Row)
		maxX = max(maxX, point.Col)
		maxY = max(maxY, point.Row)
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if ContainsPosition(points, Position{x, y}) {
				print("X")
			} else {
				print(".")
			}
		}
		println()
	}
}

// move Moves a position in a direction
func (p Position) move(direction Direction, distance int) Position {
	switch direction {
	case Up:
		return Position{p.Col, p.Row + distance}
	case Right:
		return Position{p.Col + distance, p.Row}
	case Down:
		return Position{p.Col, p.Row - distance}
	case Left:
		return Position{p.Col - distance, p.Row}
	}
	panic("invalid direction")
}
