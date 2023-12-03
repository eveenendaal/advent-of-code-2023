package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type Point struct {
	x     int
	y     int
	value string
}

type Number struct {
	value  int
	points []Point
}

// Part1 is the solution to day3 part 1
func Part1(filePath string) {
	// Open File
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 0
	total := 0
	symbols := make([]Point, 0)
	numberPoints := make([]Point, 0)
	symbolChars := []string{"#", "*", "$", "+"}

	// Scan file and finding symbols
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Find Symbols
		for x, c := range line {
			// If a symbol is found, add it to the map
			for _, symbolChar := range symbolChars {
				if string(c) == symbolChar {
					symbols = append(symbols, Point{x: x, y: y})
				}
			}

		}

		for x, c := range line {
			// If a "." is found reset the numberPoints and store the number unless there are no points
			if unicode.IsDigit(c) == false {
				if len(numberPoints) > 0 {
					// Concat the numberPoints into a number
					numberString := ""
					for _, point := range numberPoints {
						numberString = numberString + point.value
					}
					// convert numberString to an int
					number, _ := strconv.Atoi(numberString)
					if checkDistance(numberPoints, symbols) {
						// Add number to total
						total += number
					}

					// Clear numberPoints
					numberPoints = make([]Point, 0)
				}
			} else {
				// add point to numberPoints
				numberPoints = append(numberPoints, Point{x: x, y: y, value: string(c)})
			}

		}

		y = y + 1
	}

	// Print out symbols
	for _, symbol := range symbols {
		fmt.Printf("Symbol: %v\n", symbol)
	}

	// Print total
	fmt.Printf("Total: %v\n", total)
}

func checkDistance(points []Point, symbols []Point) bool {
	touching := false
	// Iterate through points
	for _, point := range points {
		// Iterate through symbols
		for _, symbol := range symbols {
			// Calculate distance
			distanceX := math.Abs(float64(point.x) - float64(symbol.x))
			distanceY := math.Abs(float64(point.y) - float64(symbol.y))
			// If distance is 1 or less, set touching to true
			if (distanceX+distanceY) == 1 || (distanceX == 1 && distanceY == 1) {
				touching = true
			}
		}
	}
	return touching

}

func main() {
	filePath := "data.txt"
	Part1(filePath)

}
