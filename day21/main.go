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

func getRelativePosition(p aoc.Position, maxX int, maxY int) aoc.Position {
	newCol := p.Col % maxX
	newRow := p.Row % maxY
	if newCol < 0 {
		newCol += maxX
	}
	if newRow < 0 {
		newRow += maxY
	}
	return aoc.Position{Col: newCol, Row: newRow}
}

var directions = []aoc.Position{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func findNeighbors(position aoc.Position, positions []aoc.Position, maxX, maxY int) []aoc.Position {
	neighbors := make([]aoc.Position, 0)
	for _, direction := range directions {
		neighbor := aoc.Position{Col: position.Col + direction.Col, Row: position.Row + direction.Row}
		relativeNeighborPosition := getRelativePosition(neighbor, maxX, maxY)
		if aoc.ContainsPosition(positions, relativeNeighborPosition) && !aoc.ContainsPosition(neighbors, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func findMaxes(gardens []aoc.Position) (int, int) {
	maxX := 0
	maxY := 0
	for _, garden := range gardens {
		nextX := garden.Col + 1
		nextY := garden.Row + 1
		if nextY > maxY {
			maxY = nextY
		}
		if nextX > maxX {
			maxX = nextX
		}
	}
	return maxX, maxY
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

func Part1(filePath string, totalSteps int) int64 {
	start, gardens := parseInput(filePath)

	var visited = make(map[int][]aoc.Position)
	visited[0] = append(visited[0], start)
	//found := 0
	//previousLen := int64(0)
	maxX, maxY := findMaxes(gardens)

	for move := 0; move < totalSteps; move++ {
		for _, currentPos := range visited[move] {
			neighbors := findNeighbors(currentPos, gardens, maxX, maxY)
			for _, neighbor := range neighbors {
				if !aoc.ContainsPosition(visited[move+1], neighbor) {
					visited[move+1] = append(visited[move+1], neighbor)
				}
			}
		}
		// fmt.Printf("Move %d: %d\n", move, len(visited[move]))
	}

	sum := int64(len(visited[len(visited)-1]))
	return sum
}

func f(x int64, a [3]int64) int64 {
	b0 := a[0]
	b1 := a[1] - a[0]
	b2 := a[2] - a[1]
	return b0 + b1*x + (x*(x-1)/2)*(b2-b1)
}

func Part2(filePath string, totalSteps int) int64 {
	start, gardens := parseInput(filePath)
	var a [3]int64
	found := 0

	var visited = make(map[int][]aoc.Position)
	visited[0] = append(visited[0], start)
	maxX, maxY := findMaxes(gardens)
	prevLen := int64(0)

	for move := 0; move < totalSteps; move++ {
		for _, currentPos := range visited[move] {
			neighbors := findNeighbors(currentPos, gardens, maxX, maxY)
			for _, neighbor := range neighbors {
				if !aoc.ContainsPosition(visited[move+1], neighbor) {
					visited[move+1] = append(visited[move+1], neighbor)
				}
			}
		}

		if (move % maxX) == (totalSteps % maxX) {
			fmt.Println("Move", move, len(visited[move]), "prevLen", int64(len(visited[move]))-prevLen)
			prevLen = int64(len(visited[move]))
			a[found] = prevLen
			found++
		}
		if found == 3 {
			break
		}
	}

	sum := f(int64(totalSteps/maxX), a)
	return sum
}
