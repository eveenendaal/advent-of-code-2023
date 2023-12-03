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
	symbols := make([]Point, 0)
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

		y = y + 1
	}

	// Print out symbols
	for _, symbol := range symbols {
		fmt.Printf("Symbol: %v\n", symbol)
	}

}

func main() {
	filePath := "data.txt"
	Part1(filePath)

}
