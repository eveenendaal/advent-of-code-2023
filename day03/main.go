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

func findSymbols(filePath string) []Point {
	// Open File
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 0
	symbols := make([]Point, 0)

	// Scan file and finding symbols
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Find Symbols
		for x, c := range line {
			// If a symbol is found, add it to the map
			if unicode.IsDigit(c) == false && string(c) != "." {
				symbols = append(symbols, Point{x: x, y: y, value: string(c)})
			}
		}

		y = y + 1
	}

	return symbols
}

// Part1 is the solution to day3 part 1
func Part1(filePath string) {
	symbols := findSymbols(filePath)

	// Open File
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 0
	total := 0
	numberPoints := make([]Point, 0)

	// Scan file and finding numbers
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

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

	// Print total
	fmt.Printf("Part 1 Total: %v\n", total)
}

func findGears(filePath string) []Point {
	symbols := findSymbols(filePath)
	gears := make([]Point, 0)
	// Filter to just "*" symbols
	for _, symbol := range symbols {
		if symbol.value == "*" {
			// Find points that are touching the symbol
			gears = append(gears, symbol)
		}
	}
	return gears
}

func findNumbersByGear(gear Point, filePath string) []int {
	// Open File
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numberPoints := make([]Point, 0)
	numbers := make([]int, 0)
	y := 0

	// gear into a slice
	gears := make([]Point, 0)
	gears = append(gears, gear)

	// Scan file and finding numbers
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

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
					if checkDistance(numberPoints, gears) {
						// Add number to total
						numbers = append(numbers, number)
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

	fmt.Printf("Numbers: %v\n", numbers)
	return numbers
}

func Part2(filePath string) {
	gears := findGears(filePath)
	total := 0
	for _, gear := range gears {
		numbers := findNumbersByGear(gear, filePath)
		if len(numbers) == 2 { // If there are two numbers, multiply them
			total += numbers[0] * numbers[1]
		}
	}

	// Print symbols
	fmt.Printf("Gears: %v\n", gears)
	fmt.Printf("Part 2 Total: %v\n", total)
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
			distance := int(math.Max(distanceX, distanceY))

			// Print points and distance
			// fmt.Printf("Point: %v, Symbol: %v, Distance: %v\n", point, symbol, distance)

			// If distance is 1 or less, set touching to true
			if distance <= 1 {
				touching = true
			}
		}
	}

	// // If not touching print points
	// if touching == false {
	// 	fmt.Printf("Points: %v\n", points)
	// }

	return touching

}

func main() {
	filePath := "data.txt"
	//Part1(filePath)
	Part2(filePath)

}
