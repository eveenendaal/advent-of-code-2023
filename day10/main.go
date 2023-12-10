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

func directionsContainPoint(option Node, point Point) bool {
	if len(option.directions) == 0 {
		return false
	}

	for _, p := range option.directions {
		next := Point{option.point.x + p.x, option.point.y + p.y}
		if next == point {
			return true
		}
	}
	return false
}

func findDirection(vertical []Point, point Point) bool {
	for _, p := range vertical {
		if p == point {
			return true
		}
	}
	return false
}

func Part2(filePath string) int {

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
				points[point] = Node{char, point, []Point{}}
			case '|':
				points[point] = Node{char, point, []Point{{0, -1}, {0, 1}}}
			case '-':
				points[point] = Node{char, point, []Point{{-1, 0}, {1, 0}}}
			case 'J':
				points[point] = Node{char, point, []Point{{-1, 0}, {0, -1}}}
			case 'L':
				points[point] = Node{char, point, []Point{{1, 0}, {0, -1}}}
			case 'F':
				points[point] = Node{char, point, []Point{{1, 0}, {0, 1}}}
			case '7':
				points[point] = Node{char, point, []Point{{-1, 0}, {0, 1}}}
			case '.':
				points[point] = Node{char, point, []Point{}}
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
	if directionsContainPoint(option1, start) {
		paths = append(paths, option1.point)
	}
	if directionsContainPoint(option2, start) {
		paths = append(paths, option2.point)
	}
	if directionsContainPoint(option3, start) {
		paths = append(paths, option3.point)
	}
	if directionsContainPoint(option4, start) {
		paths = append(paths, option4.point)
	}

	history := []Point{start}
	fmt.Printf("Start: %v\n", start)
	// fmt.Printf("Points: %v\n", points)
	fmt.Printf("Paths: %v\n", paths)
	nextNode := points[paths[0]]
	for {
		current := nextNode.point
		nextDirections := nextNode.directions
		// fmt.Printf("Current: %v - %v\n", current, string(nextNode.character))
		// Filter out directions in the history
		if len(nextDirections) == 0 {
			// Throw error
			fmt.Printf("No directions for %v\n", nextNode)
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

		if len(history) > 1 && (pointOne == start || pointTwo == start) {
			break
		} else if !foundOne {
			nextNode = points[pointOne]
			history = append(history, current)
		} else if !foundTwo {
			nextNode = points[pointTwo]
			history = append(history, current)
		} else {
			fmt.Printf("History: %v\n", history)
			fmt.Printf("No valid directions for %v\n", nextNode)
			panic("No valid directions for current point")
		}
	}
	history = append(history, start)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(history); i++ {
		cur := history[i]
		next := history[(i+1)%len(history)]

		polygonArea += cur.x*next.y - cur.y*next.x
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return -1
	}

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return polygonArea - len(history)/2 + 1
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
				points[point] = Node{char, point, []Point{{1, 0}, {0, -1}}}
			case 'F':
				points[point] = Node{char, point, []Point{{1, 0}, {0, 1}}}
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
	if directionsContainPoint(option1, start) {
		paths = append(paths, option1.point)
	}
	if directionsContainPoint(option2, start) {
		paths = append(paths, option2.point)
	}
	if directionsContainPoint(option3, start) {
		paths = append(paths, option3.point)
	}
	if directionsContainPoint(option4, start) {
		paths = append(paths, option4.point)
	}

	history := []Point{start}
	fmt.Printf("Start: %v\n", start)
	// fmt.Printf("Points: %v\n", points)
	fmt.Printf("Paths: %v\n", paths)
	nextNode := points[paths[0]]
	for {
		current := nextNode.point
		nextDirections := nextNode.directions
		// fmt.Printf("Current: %v - %v\n", current, string(nextNode.character))
		// Filter out directions in the history
		if len(nextDirections) == 0 {
			// Throw error
			fmt.Printf("No directions for %v\n", nextNode)
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

		if len(history) > 1 && (pointOne == start || pointTwo == start) {
			break
		} else if !foundOne {
			nextNode = points[pointOne]
			history = append(history, current)
		} else if !foundTwo {
			nextNode = points[pointTwo]
			history = append(history, current)
		} else {
			fmt.Printf("History: %v\n", history)
			fmt.Printf("No valid directions for %v\n", nextNode)
			panic("No valid directions for current point")
		}
	}

	// fmt.Printf("Points: %v\n", points)
	total := 0
	fmt.Printf("Start: %v, Total: %d\n", start, total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return -1
	}

	return total
}

func main() {
	// fmt.Println("Part 1:", Part1("data.txt"))
	fmt.Println("Part 2:", Part2("data.txt"))
}
