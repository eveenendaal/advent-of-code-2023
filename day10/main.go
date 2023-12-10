package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type Node struct {
	character  rune
	point      Point
	directions []Point
}

func checkHistory(history []Point, point Point) bool {
	for _, p := range history {
		if p == point {
			return true
		}
	}
	return false
}

func Part1(filePath string) int {

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	y := 0
	start := Point{0, 0}
	// Make map of points
	points := make(map[Point]Node)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Read each character in the line
		for x, char := range line {
			point := Point{x, y}
			switch char {
			case 'S':
				start = point
			case '|':
				points[point] = Node{char, point, []Point{{0, -1}, {0, 1}}}
			case '-':
				points[point] = Node{char, point, []Point{{-1, 0}, {1, 0}}}
			case 'J':
				points[point] = Node{char, point, []Point{{-1, 0}, {0, -1}}}
			case 'L':
				points[point] = Node{char, point, []Point{{1, 0}, {0, 1}}}
			case 'F':
				points[point] = Node{char, point, []Point{{1, 0}, {0, -1}}}
			case '7':
				points[point] = Node{char, point, []Point{{-1, 0}, {0, 1}}}
			case '.':
				break
			default:
				fmt.Println("Unknown character:", string(char))
			}
		}

		y++
	}

	// Check the 4 points around the start
	option1 := points[Point{start.x, start.y - 1}]
	option2 := points[Point{start.x, start.y + 1}]
	option3 := points[Point{start.x - 1, start.y}]
	option4 := points[Point{start.x + 1, start.y}]

	paths := []Point{}
	if option1.directions != nil {
		paths = append(paths, option1.point)
	}
	if option2.directions != nil {
		paths = append(paths, option2.point)
	}
	if option3.directions != nil {
		paths = append(paths, option3.point)
	}
	if option4.directions != nil {
		paths = append(paths, option4.point)
	}

	history := []Point{start}
	done := false
	stepCount := 0
	fmt.Printf("Points: %v\n", points)
	for {
		for i, current := range paths {
			nextNode := points[current]
			nextDirections := nextNode.directions
			fmt.Printf("Current: %v - %v (%d)\n", current, string(nextNode.character), i)
			// Filter out directions in the history
			if len(nextDirections) == 0 {
				// Throw error
				fmt.Printf("No directions for %v\n", current)
				panic("No directions for current point")
			}

			// Check for direction one
			directionOne := nextDirections[0]
			pointOne := Point{current.x + directionOne.x, current.y + directionOne.y}
			foundOne := checkHistory(history, pointOne)

			// Check for direction two
			directionTwo := nextDirections[1]
			pointTwo := Point{current.x + directionTwo.x, current.y + directionTwo.y}
			foundTwo := checkHistory(history, pointTwo)

			if foundOne && foundTwo {
				fmt.Printf("Done: %v\n", current)
				done = true
				break
			} else if !foundOne {
				paths[i] = pointOne
				history = append(history, current)
			} else if !foundTwo {
				paths[i] = pointTwo
				history = append(history, current)
			}
		}

		stepCount++
		if done {
			break
		}
	}

	fmt.Printf("Points: %v\n", points)
	fmt.Printf("Start: %v, Total: %d\n", start, stepCount)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return -1
	}

	return stepCount
}

func main() {
	Part1("data.txt")
}
