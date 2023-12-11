package main

import (
	"bufio"
	"fmt"
	"os"
)

type Star struct {
	x      int
	y      int
	number int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s Star) distance(other Star) int {
	return abs(s.x-other.x) + abs(s.y-other.y)
}

func (s Star) finalLocation(emptyX []int, emptyY []int) Star {
	finalX := s.x
	finalY := s.y
	// filter emptyY less then x
	for _, x := range emptyX {
		if x < s.x {
			finalX++
		}
	}
	// filter emptyX less then y
	for _, y := range emptyY {
		if y < s.y {
			finalY++
		}
	}
	return Star{finalX, finalY, s.number}
}

func contains(stars []Star, star Star) bool {
	for _, s := range stars {
		if s.x == star.x && s.y == star.y {
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
		panic(err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	stars := []Star{}

	// Read the file line by line
	y := 0
	starCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Read each character in the line
		for x, char := range line {
			// Process each character here
			if string(char) == "#" {
				stars = append(stars, Star{x, y, starCount})
				starCount++
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	// Find the max X and Y
	maxX := 0
	maxY := 0
	for _, star := range stars {
		if star.x > maxX {
			maxX = star.x
		}
		if star.y > maxY {
			maxY = star.y
		}
	}

	emptyX := []int{}
	emptyY := []int{}

	for y := 0; y < maxY; y++ {
		empty := true
		for _, star := range stars {
			if star.y == y {
				empty = false
				break
			}
		}
		if empty {
			emptyY = append(emptyY, y)
		}
	}
	for x := 0; x < maxX; x++ {
		empty := true
		for _, star := range stars {
			if star.x == x {
				empty = false
				break
			}
		}
		if empty {
			emptyX = append(emptyX, x)
		}
	}

	// Expand the grid
	realStars := []Star{}
	for _, star := range stars {
		realStars = append(realStars, star.finalLocation(emptyX, emptyY))
	}

	fmt.Printf("Real Stars: %v\n", realStars)

	total := 0

	// Iterate over every combination of stars
	handledStars := []Star{}
	for _, star1 := range realStars {
		handledStars = append(handledStars, star1)
		for _, star2 := range realStars {
			if contains(handledStars, star2) {
				continue
			}
			distance := star1.distance(star2)
			// fmt.Printf("Star1: %v, Star2: %v, Distance %d\n", star1, star2, distance)
			total = distance + total
		}
	}

	return total
}

func main() {
	fmt.Println("Part 1:", Part1("data.txt"))
}
